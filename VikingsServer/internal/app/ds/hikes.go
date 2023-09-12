package ds

type Hikes struct {
	ID          uint   `json:"id"`
	HikeName    string `json:"hikeName"`
	DateStart   string `json:"dateStart"`
	DateEnd     string `json:"dateEnd"`
	Leader      string `json:"leader"`
	Status      string `json:"status"`
	Description string `json:"description"`
}
