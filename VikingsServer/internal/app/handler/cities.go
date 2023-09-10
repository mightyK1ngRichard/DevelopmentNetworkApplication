package handler

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/utils"
	"fmt"
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
	data := ds.CityViewData{}
	citiesList, err := h.Repository.CitiesList()
	if err != nil {
		// TODO: Обработать ошибку.
		return
	}
	data.Cities = citiesList
	searchText := ctx.Query("search")
	if idStr := ctx.Query("city"); idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return
		}
		var currentCity ds.City
		for _, city := range *data.Cities {
			if uint(id) == city.ID {
				currentCity = city
				break
			}
		}

		if currentCity != (ds.City{}) {
			index := utils.FindElement(*data.Cities, currentCity)
			var lookAlso []ds.City
			if index != -1 {
				startIndex := index + 1
				if startIndex >= len(*data.Cities) {
					startIndex = 0
				}
				endIndex := utils.Min(startIndex+5, len(*data.Cities))
				if startIndex < endIndex {
					lookAlso = (*data.Cities)[startIndex:endIndex]
				} else {
					lookAlso = (*data.Cities)[endIndex:startIndex]
				}
			}

			ctx.HTML(http.StatusOK, "city.card.tmpl",
				ds.CityViewData{
					Cities:   &[]ds.City{currentCity},
					LookAlso: &lookAlso,
				},
			)
		}

		return
	}

	if searchText != "" {
		var filteredCities []ds.City
		for _, city := range *data.Cities {
			if strings.Contains(strings.ToLower(city.CityName), strings.ToLower(searchText)) {
				filteredCities = append(filteredCities, city)
			}
		}
		ctx.HTML(http.StatusOK, "cities.tmpl", ds.CityViewData{Cities: &filteredCities})

		return
	}

	ctx.HTML(http.StatusOK, "cities.tmpl", data)
}

func (h *Handler) CitiesDelete(ctx *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repository.DeleteCity(request.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "couldn't delete the city",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": fmt.Sprintf("deleted city with id: %d", request.ID),
	})
}

func (h *Handler) DeleteCityWithStatus(ctx *gin.Context) {
	id := ctx.PostForm("cityID")

	if err := h.Repository.DeleteCityWithStatus(id); err != nil {
		h.Logger.Error("couldn't delete city")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete city"})
		return
	}
	h.Logger.Info("city with id=" + id + "update success")
	ctx.Redirect(http.StatusSeeOther, citiesHTML+"?city="+id)
}
