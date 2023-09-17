package ds

import (
	"time"
)

type Hike struct {
	ID                 uint       `json:"id" gorm:"primary_key"`
	HikeName           string     `json:"hike_name" gorm:"type:varchar(50);not null"`
	DateStart          time.Time  `json:"date_start" gorm:"type:date;not null;default:current_date"`
	DateEnd            time.Time  `json:"date_end" gorm:"type:date"`
	DateStartPreparing time.Time  `json:"date_start_preparing" gorm:"type:date"`
	AuthorID           uint       `json:"author_id"`
	Author             Author     `json:"author" gorm:"foreignkey:AuthorID"`
	StatusID           uint       `json:"status_id"`
	Status             HikeStatus `json:"status" gorm:"foreignkey:StatusID"`
	Description        string     `json:"description" gorm:"type:text"`
}
