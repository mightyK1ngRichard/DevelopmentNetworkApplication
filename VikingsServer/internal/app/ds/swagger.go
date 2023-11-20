package ds

type DeleteCityRes struct {
	DeletedId int `json:"deleted_id"`
}

type DeleteCityReq struct {
	ID string `json:"id"`
}

type UpdateCityReq struct {
	Id          int    `json:"id" binding:"required"`
	CityName    string `json:"city_name"`
	Description string `json:"description"`
	StatusId    int    `json:"status_id"`
}

type AddCityIntoHikeReq struct {
	CityID       int `json:"city_id" binding:"required" example:"1"`
	SerialNumber int `json:"serial_number" binding:"required" example:"1"`
}

type AddCityIntoHikeResp struct {
	Status string `json:"status"`
	Id     int    `json:"id"`
}

type UpdateCityResp struct {
	ID          string `json:"id"`
	CityName    string `json:"city_name"`
	StatusID    string `json:"status_id"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type AddCityResp struct {
	Status string `json:"status"`
	CityId string `json:"city_id"`
}

type CitiesListResp struct {
	Status   string `json:"status"`
	Cities   []City `json:"cities"`
	BasketId string `json:"basket_id"`
}

type AddCityIntoHikeRequest struct {
}
