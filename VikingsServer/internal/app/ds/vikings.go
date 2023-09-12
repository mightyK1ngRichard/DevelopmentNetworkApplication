package ds

type Vikings struct {
	ID          int    `json:"id"`
	VikingName  string `json:"vikingName"`
	Post        string `json:"post"`
	Birthday    string `json:"birthday,omitempty"`
	DayOfDeath  string `json:"dayOfDeath,omitempty"`
	CityOfBirth int    `json:"cityOfBirth,omitempty"`
	ImageURL    string `json:"imageURL"`
}
