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
	cities, err := h.Repository.CitiesList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"cities": cities,
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

		var lookAlso []ds.City
		if currentCity != (ds.City{}) {
			index := utils.FindElement(*data.Cities, currentCity)
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
		}

		ctx.HTML(http.StatusOK, "city.card.tmpl",
			ds.OneCityViewData{
				City:     &currentCity,
				LookAlso: &lookAlso,
			},
		)

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

func (h *Handler) CitiesDeleteCascade(ctx *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repository.DeleteCity(request.ID); err != nil {
		h.Logger.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "couldn't delete the city: " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": fmt.Sprintf("deleted city with id: %d", request.ID),
	})
}

func (h *Handler) DeleteCityWithStatus(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("cityID"))
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if err := h.Repository.DeleteCityWithStatus(id); err != nil {
		h.Logger.Error("couldn't delete city")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete city"})
		return
	}
	h.Logger.Info("city with id=" + "update success")
	ctx.Redirect(http.StatusSeeOther, citiesHTML)
}

func (h *Handler) DeleteCity(ctx *gin.Context) {
	var requestData struct {
		ID string `json:"id"`
	}
	if err := ctx.BindJSON(&requestData); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	id, errInt := strconv.Atoi(requestData.ID)
	if errInt != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, errInt)
		return
	}
	if err := h.Repository.DeleteCityWithStatus(id); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.Logger.Info("city with id=" + fmt.Sprintf("%d", id) + "update success")
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "cityId": id})
}
