package ds

type Hikes struct {
	ID          int    `json:"id"`
	HikeName    string `json:"hikeName"`
	DateStart   string `json:"dateStart"`
	DateEnd     string `json:"dateEnd"`
	Leader      int    `json:"leader"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}
