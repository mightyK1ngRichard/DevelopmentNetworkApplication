package handler

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

// /Users/dmitriy/go/bin/swag init -g cmd/main/main.go

// CitiesList godoc
// @Summary Get a list of cities
// @Description Get a list of cities with optional filtering by city name.
// @Tags cities
// @Produce json
// @Param city query string false "City name for filtering"
// @Success 200 {array} ds.City
// @Router /cities [get]
func (h *Handler) CitiesList(ctx *gin.Context) {
	if idStr := ctx.Query("city"); idStr != "" {
		cityById(ctx, h, idStr)
		return
	}

	cities, err := h.Repository.CitiesList()
	if err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	searchText := ctx.Query("search")
	if searchText != "" {
		var filteredCities []ds.City
		for _, city := range *cities {
			if strings.Contains(strings.ToLower(city.CityName), strings.ToLower(searchText)) {
				filteredCities = append(filteredCities, city)
			}
		}
		h.successHandler(ctx, "cities", filteredCities)
		return
	}

	h.successHandler(ctx, "cities", cities)
}

func cityById(ctx *gin.Context, h *Handler, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	city, errBD := h.Repository.CitiesById(uint(id))
	if errBD != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, errBD)
		return
	}

	h.successHandler(ctx, "city", city)
}

func (h *Handler) DeleteCityWithParam(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err2 := strconv.Atoi(idStr)
	if err2 != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err2)
		return
	}
	if id == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.DeleteCity(uint(id)); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "deleted_id", id)
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
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.DeleteCity(uint(id)); err != nil {
		h.errorHandler(ctx, http.StatusInternalServerError, err)
		return
	}

	h.successHandler(ctx, "deleted_id", id)
}

func (h *Handler) AddImage(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	cityID := ctx.Request.FormValue("city_id")

	if cityID == "" {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if header == nil || header.Size == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, headerNotFound)
		return
	}
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	defer func(file multipart.File) {
		errLol := file.Close()
		if errLol != nil {
			h.errorHandler(ctx, http.StatusInternalServerError, errLol)
			return
		}
	}(file)

	// Upload the image to minio server.
	newImageURL, errorCode, errImage := h.createImageCity(&file, header, cityID)
	if errImage != nil {
		h.errorHandler(ctx, errorCode, errImage)
		return
	}

	h.successAddHandler(ctx, "image_url", newImageURL)
}

// Функция записи фото в минио
func (h *Handler) createImageCity(
	file *multipart.File,
	header *multipart.FileHeader,
	cityID string,
) (string, int, error) {
	newImageURL, errMinio := h.createImageInMinio(file, header)
	if errMinio != nil {
		return "", http.StatusInternalServerError, errMinio
	}
	if err := h.Repository.UpdateCityImage(cityID, newImageURL); err != nil {
		return "", http.StatusInternalServerError, err
	}
	return newImageURL, 0, nil
}

func (h *Handler) AddCity(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image_url")
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	cityName := ctx.Request.FormValue("city_name")
	statusID := ctx.Request.FormValue("status_id")
	description := ctx.Request.FormValue("description")
	intStatus, errStatusInt := strconv.Atoi(statusID)
	if errStatusInt != nil {
		h.errorHandler(ctx, http.StatusBadRequest, errStatusInt)
		return
	}

	newCity := ds.City{
		CityName:    cityName,
		StatusID:    uint(intStatus),
		Description: description,
	}
	if errorCode, errCreate := h.createCity(&newCity); err != nil {
		h.errorHandler(ctx, errorCode, errCreate)
		return
	}
	newImageURL, errCode, errDB := h.createImageCity(&file, header, fmt.Sprintf("%d", newCity.ID))
	if errDB != nil {
		h.errorHandler(ctx, errCode, errDB)
		return
	}
	newCity.ImageURL = newImageURL

	h.successAddHandler(ctx, "city_id", newCity.ID)
}

func (h *Handler) createCity(city *ds.City) (int, error) {
	if city.ID != 0 {
		return http.StatusBadRequest, idMustBeEmpty
	}
	if city.CityName == "" {
		return http.StatusBadRequest, cityCannotBeEmpty
	}
	if err := h.Repository.AddCity(city); err != nil {
		return http.StatusBadRequest, err
	}
	return 0, nil
}

func (h *Handler) UpdateCity(ctx *gin.Context) {
	var updatedCity ds.City
	if err := ctx.BindJSON(&updatedCity); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if updatedCity.ID == 0 {
		h.errorHandler(ctx, http.StatusBadRequest, idNotFound)
		return
	}
	if err := h.Repository.UpdateCity(&updatedCity); err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}

	h.successHandler(ctx, "updated_city", gin.H{
		"id":          updatedCity.ID,
		"city_name":   updatedCity.CityName,
		"status_id":   updatedCity.StatusID,
		"description": updatedCity.Description,
		"image_url":   updatedCity.ImageURL,
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
		cityByIdHTML(ctx, &data, idStr)
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

func cityByIdHTML(ctx *gin.Context, data *ds.CityViewData, idStr string) {
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
}

func (h *Handler) DeleteCityHTML(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.PostForm("cityID"))
	if err != nil {
		h.errorHandler(ctx, http.StatusBadRequest, err)
		return
	}
	if err := h.Repository.DeleteCity(uint(id)); err != nil {
		h.Logger.Error("couldn't delete city")
		ctx.Redirect(http.StatusSeeOther, citiesHTML)
		return
	}
	ctx.Redirect(http.StatusSeeOther, citiesHTML)
}
