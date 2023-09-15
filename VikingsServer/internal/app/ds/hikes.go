package ds

type Hikes struct {
	ID          uint   `json:"id"`
	HikeName    string `json:"hikeName"`
	DateStart   string `json:"dateStart"`
	DateEnd     string `json:"dateEnd"`
	AuthorID    int    `json:"authorId"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}
