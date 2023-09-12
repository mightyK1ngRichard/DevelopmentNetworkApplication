package handler

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (h *Handler) CitiesList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"cities": "cities",
	})
}

func (h *Handler) CitiesHTML(ctx *gin.Context) {
	data := ds.GetCityViewData()
	searchText := ctx.Query("search")

	if idStr := ctx.Query("city"); idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return
		}
		var currentCity ds.City
		for _, city := range data.Cities {
			if uint(id) == city.ID {
				currentCity = city
				break
			}
		}

		if currentCity != (ds.City{}) {
			index := utils.FindElement(data.Cities, currentCity)
			var lookAlso []ds.City
			if index != -1 {
				startIndex := index + 1
				if startIndex >= len(data.Cities) {
					startIndex = 0
				}
				endIndex := utils.Min(startIndex+5, len(data.Cities))
				if startIndex < endIndex {
					lookAlso = data.Cities[startIndex:endIndex]
				} else {
					lookAlso = data.Cities[endIndex:startIndex]
				}
			}

			ctx.HTML(http.StatusOK, "city.card.tmpl",
				ds.CityViewData{
					Cities:   []ds.City{currentCity},
					LookAlso: lookAlso,
				},
			)
		}
		return
	}

	if searchText != "" {
		var filteredCities []ds.City
		for _, city := range data.Cities {
			if strings.Contains(strings.ToLower(city.Name), strings.ToLower(searchText)) {
				filteredCities = append(filteredCities, city)
			}
		}
		ctx.HTML(http.StatusOK, "cities.tmpl", ds.CityViewData{Cities: filteredCities})
		return
	}

	ctx.HTML(http.StatusOK, "cities.tmpl", data)
}
