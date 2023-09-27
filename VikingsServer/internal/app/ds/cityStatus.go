package ds

type CityStatus struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	StatusName string `json:"statusName" gorm:"type:varchar(30);not null;->"`
}
