package ds

type Vikings struct {
	ID          int    `json:"id"`
	VikingName  string `json:"vikingName"`
	Post        string `json:"post"`
	Birthday    string `json:"birthday"`
	DayOfDeath  string `json:"dayOfDeath"`
	CityOfBirth int    `json:"cityOfBirth"`
	ImageURL    string `json:"imageURL"`
}
