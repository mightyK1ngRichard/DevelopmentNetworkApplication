package ds

type HikeStatus struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	StatusName string `json:"statusName" gorm:"->"`
}
