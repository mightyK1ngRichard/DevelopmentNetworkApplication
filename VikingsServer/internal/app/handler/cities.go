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
		h.errorHandler(ctx, http.StatusInternalServerError, err)
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
		h.errorHandler(ctx, http.StatusInternalServerError, err)
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
		ID uint `json:"id"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repository.DeleteCity(request.ID); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": fmt.Sprintf("deleted city with id: %d", request.ID),
	})
}

func (h *Handler) DeleteCityHTML(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("cityID"))
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if err := h.Repository.DeleteCity(uint(id)); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
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
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	id, err2 := strconv.Atoi(request.ID)
	if err2 != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err2)
		return
	}
	if id == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, fmt.Errorf("incorrect id"))
		return
	}
	if err := h.Repository.DeleteCity(uint(id)); err != nil {
		h.Logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Redirect(http.StatusSeeOther, citiesHTML)
}
