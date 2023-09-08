package handler

import (
	"05-lab-01-DimaPermyakov/internal/app/ds"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CitiesList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"cities": "cities",
	})
}

func (h *Handler) CitiesHTML(ctx *gin.Context) {
	data := ds.CityViewData{
		Cities: []ds.City{
			{ID: 0, Name: "Dima 1", ImageURL: "http://localhost:7070/static/img/image1.jpeg"},
			{ID: 1, Name: "Dima 2", ImageURL: "http://localhost:7070/static/img/image2.jpg"},
			{ID: 2, Name: "Dima 3", ImageURL: "http://localhost:7070/static/img/image3.jpg"},
			{ID: 3, Name: "Dima 4", ImageURL: "http://localhost:7070/static/img/image4.jpg"},
			{ID: 4, Name: "Dima 5", ImageURL: "http://localhost:7070/static/img/image5.jpg"},
			{ID: 5, Name: "Dima 6", ImageURL: "http://localhost:7070/static/img/image6.jpg"},
			{ID: 6, Name: "Dima 7", ImageURL: "http://localhost:7070/static/img/image7.jpg"},
		},
	}

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
			ctx.HTML(http.StatusOK, "city.card.tmpl", currentCity)
		}
		return
	}

	ctx.HTML(http.StatusOK, "cities.tmpl", data)
}
