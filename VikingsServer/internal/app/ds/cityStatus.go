package ds

import "gorm.io/gorm"

type CityStatus struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primary_key"`
	StatusName string `json:"statusName" gorm:"type:varchar(30);not null;->"`
}
