package ds

type Author struct {
	ID         int    `json:"id"`
	AuthorName string `json:"authorName"`
	Profession string `json:"profession"`
	Birthday   string `json:"birthday"`
	ImageURL   string `json:"imageURL"`
}
