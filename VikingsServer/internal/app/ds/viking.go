package ds

import (
	"time"
)

type Viking struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	VikingName    string    `json:"viking_name" gorm:"type:varchar(60);not null"`
	Post          string    `json:"post" gorm:"type:varchar(100);not null"`
	Birthday      time.Time `json:"birthday"`
	DayOfDeath    time.Time `json:"day_of_death"`
	CityOfBirthID uint      `json:"city_of_birth_id"`
	CityOfBirth   City      `json:"city_of_birth" gorm:"foreignkey:CityOfBirthID"`
	ImageURL      string    `json:"image_url" gorm:"type:varchar(500);default:'https://novye-multiki.ru/wp-content/uploads/2019/01/kak-priruchit-drakona-3-oboi8.jpg'"`
}
