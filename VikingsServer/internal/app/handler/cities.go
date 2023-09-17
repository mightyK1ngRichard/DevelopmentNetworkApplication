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

func (h *Handler) DeleteCityHTML(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("cityID"))
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if err := h.Repository.DeleteCity(uint(id)); err != nil {
		h.Logger.Error("couldn't delete city")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete city"})
		return
	}
	h.Logger.Info("city with id=" + "update success")
	ctx.Redirect(http.StatusSeeOther, citiesHTML)
}

func (h *Handler) DeleteCity(ctx *gin.Context) {
	var request struct {
		ID string `json:"id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		h.Logger.Error(err)
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	id, err2 := strconv.Atoi(request.ID)
	if err2 != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err2)
		return
	}
	if id == 0 {
		h.Logger.Error("incorrect id")
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("incorrect id"))
		return
	}
	if err := h.Repository.DeleteCity(uint(id)); err != nil {
		h.Logger.Error(err)
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"deleted_id": id,
	})
}

func (h *Handler) AddCity(ctx *gin.Context) {
	var newCity ds.City
	if err := ctx.BindJSON(&newCity); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if newCity.ID != 0 {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("id must be empty"))
		return
	}
	if newCity.CityName == "" {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("city name cannot be empty"))
		return
	}
	if err := h.Repository.AddCity(&newCity); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"city_id": newCity.ID,
	})
}

func (h *Handler) UpdateCity(ctx *gin.Context) {
	var updatedCity ds.City
	if err := ctx.BindJSON(&updatedCity); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if updatedCity.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("id not found"))
		return
	}
	if err := h.Repository.UpdateCity(&updatedCity); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"updated_city": gin.H{
			"id":          updatedCity.ID,
			"city_name":   updatedCity.CityName,
			"status_id":   updatedCity.StatusID,
			"description": updatedCity.Description,
			"image_url":   updatedCity.ImageURL,
		},
	})
}
