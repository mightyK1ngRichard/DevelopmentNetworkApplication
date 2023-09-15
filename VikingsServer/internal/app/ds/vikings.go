package ds

type Vikings struct {
	ID            int    `json:"id"`
	VikingName    string `json:"vikingName"`
	Post          string `json:"post"`
	Birthday      string `json:"birthday"`
	DayOfDeath    string `json:"dayOfDeath"`
	CityOfBirthID int    `json:"cityOfBirthId"`
	ImageURL      string `json:"imageURL"`
}
