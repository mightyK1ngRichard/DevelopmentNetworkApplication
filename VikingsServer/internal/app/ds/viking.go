package ds

import (
	"time"
)

type Viking struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	VikingName    string    `json:"vikingName" gorm:"type:varchar(60);not null"`
	Post          string    `json:"post" gorm:"type:varchar(100);not null"`
	Birthday      time.Time `json:"birthday"`
	DayOfDeath    time.Time `json:"dayOfDeath"`
	CityOfBirthID uint      `json:"cityOfBirthID"`
	CityOfBirth   City      `json:"cityOfBirth" gorm:"foreignkey:CityOfBirthID"`
	ImageURL      string    `json:"imageURL" gorm:"type:varchar(500);default:'https://novye-multiki.ru/wp-content/uploads/2019/01/kak-priruchit-drakona-3-oboi8.jpg'"`
}
