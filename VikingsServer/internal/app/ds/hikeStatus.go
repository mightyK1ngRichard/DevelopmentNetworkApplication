package ds

import "gorm.io/gorm"

type HikeStatus struct {
	gorm.Model
	ID         uint   `json:"id" gorm:"primary_key"`
	StatusName string `json:"statusName" gorm:"->"`
}
