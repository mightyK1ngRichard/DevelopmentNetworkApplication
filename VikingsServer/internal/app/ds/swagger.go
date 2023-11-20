package ds

type AddCityResp struct {
	Status string `json:"status"`
	cityId string `json:"city_id"`
}

type CitiesListResp struct {
	Status   string `json:"status"`
	Cities   []City `json:"cities"`
	BasketId string `json:"basket_id"`
}

type AddCityIntoHikeRequest struct {
}
