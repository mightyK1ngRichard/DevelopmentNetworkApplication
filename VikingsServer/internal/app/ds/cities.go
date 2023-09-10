package ds

type City struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
	Area        string `json:"area"`
}

type CityViewData struct {
	Cities   []City
	LookAlso []City
}

// TODO: Заменить City на этот

type NewCity struct {
	ID          uint   `json:"id"`
	CityName    string `json:"cityName"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ImageURL    string `json:"imageURL"`
}
