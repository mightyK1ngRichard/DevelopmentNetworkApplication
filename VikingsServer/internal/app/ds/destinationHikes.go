package ds

type DestinationHikes struct {
	ID     uint `json:"id" gorm:"primary_key"`
	CityID uint `json:"cityID"`
	City   City `json:"city" gorm:"foreignkey:CityID"`
	HikeID uint `json:"hikeID"`
	Hike   Hike `json:"hike" gorm:"foreignkey:HikeID"`
}
