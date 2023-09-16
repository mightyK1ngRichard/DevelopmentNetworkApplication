package ds

import (
	"gorm.io/gorm"
	"time"
)

type Author struct {
	gorm.Model
	ID         uint      `json:"id" gorm:"primary_key"`
	AuthorName string    `json:"authorName" gorm:"type:varchar(50)"`
	Profession string    `json:"profession" gorm:"type:varchar(255)"`
	Birthday   time.Time `json:"birthday"`
	ImageURL   string    `json:"imageURL" gorm:"type:varchar(500);default:'http://localhost:7070/static/img/mock-photo.png'"`
}
