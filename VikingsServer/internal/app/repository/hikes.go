package repository

import (
	"VikingsServer/internal/app/ds"
)

func (r *Repository) HikesList() (*[]ds.Hike, error) {
	var hikes []ds.Hike
	result := r.db.Preload("Participants.Viking.CityOfBirth.Status").Preload("DestinationHikes.City.Status").Preload("User").Preload("Status").Find(&hikes)
	return &hikes, result.Error
}

func (r *Repository) HikeByID(id uint) (*ds.Hike, error) {
	hike := ds.Hike{}
	result := r.db.Preload("User").Preload("Status").First(&hike, id)
	return &hike, result.Error
}

func (r *Repository) DeleteHike(id uint) error {
	hike := ds.Hike{}
	if result := r.db.First(&hike, id); result.Error != nil {
		return result.Error
	}
	var newStatus ds.HikeStatus
	if result := r.db.Where("status_name = ?", "удалён").First(&newStatus); result.Error != nil {
		return result.Error
	}
	hike.StatusID = newStatus.ID
	result := r.db.Save(&hike)
	return result.Error
}
