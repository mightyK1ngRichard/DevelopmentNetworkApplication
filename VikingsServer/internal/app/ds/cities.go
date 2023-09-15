package ds

type City struct {
	ID          uint   `json:"id"`
	CityName    string `json:"cityName"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ImageURL    string `json:"imageURL"`
}

type CityViewData struct {
	Cities   *[]City
	LookAlso *[]City
}

type OneCityViewData struct {
	City     *City
	LookAlso *[]City
}
