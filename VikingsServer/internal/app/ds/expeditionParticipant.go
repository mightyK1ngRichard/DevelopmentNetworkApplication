package ds

type ExpeditionParticipant struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	VikingID uint   `json:"vikingID"`
	Viking   Viking `json:"viking" gorm:"foreignkey:VikingID"`
	HikeID   uint   `json:"hikeID"`
	Hike     Hike   `json:"hike" gorm:"foreignkey:HikeID"`
}
