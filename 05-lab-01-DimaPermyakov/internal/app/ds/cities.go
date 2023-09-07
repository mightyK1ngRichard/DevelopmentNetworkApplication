package ds

type City struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CityViewData struct {
	Cities []City
}
