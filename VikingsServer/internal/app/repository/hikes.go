package repository

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/utils"
)

func (r *Repository) HikesList() (*[]ds.Hike, error) {
	var hikes []ds.Hike
	result := r.db.Preload("Status").Preload("DestinationHikes.City.Status").Preload("User").Preload("Status").Find(&hikes)
	return &hikes, result.Error
}

func (r *Repository) HikeByID(id uint) (*ds.Hike, error) {
	hike := ds.Hike{}
	result := r.db.Preload("User").Preload("DestinationHikes.Hike.Status").Preload("DestinationHikes.Hike.User").Preload("DestinationHikes.City.Status").Preload("Status").First(&hike, id)
	return &hike, result.Error
}

func (r *Repository) AddHike(hike *ds.Hike) error {
	result := r.db.Create(&hike)
	return result.Error
}

func (r *Repository) DeleteHike(id uint) error {
	hike := ds.Hike{}
	if result := r.db.First(&hike, id); result.Error != nil {
		return result.Error
	}
	var newStatus ds.HikeStatus
	if result := r.db.Where("status_name = ?", utils.DeletedString).First(&newStatus); result.Error != nil {
		return result.Error
	}
	hike.StatusID = newStatus.ID
	result := r.db.Save(&hike)
	return result.Error
}

func (r *Repository) UpdateHike(updatedHike *ds.Hike) error {
	oldHike := ds.Hike{}
	if result := r.db.First(&oldHike, updatedHike.ID); result.Error != nil {
		return result.Error
	}
	if updatedHike.HikeName != "" {
		oldHike.HikeName = updatedHike.HikeName
	}
	if updatedHike.DateCreated.String() != utils.EmptyDate {
		oldHike.DateCreated = updatedHike.DateCreated
	}
	if updatedHike.DateEnd.String() != utils.EmptyDate {
		oldHike.DateEnd = updatedHike.DateEnd
	}
	if updatedHike.DateStartOfProcessing.String() != utils.EmptyDate {
		oldHike.DateStartOfProcessing = updatedHike.DateStartOfProcessing
	}
	if updatedHike.DateApprove.String() != utils.EmptyDate {
		oldHike.DateApprove = updatedHike.DateApprove
	}
	if updatedHike.DateStartHike.String() != utils.EmptyDate {
		oldHike.DateStartHike = updatedHike.DateStartHike
	}
	if updatedHike.UserID != utils.EmptyInt {
		oldHike.UserID = updatedHike.UserID
	}
	if updatedHike.StatusID != utils.EmptyInt {
		oldHike.StatusID = updatedHike.StatusID
	}
	if updatedHike.Description != "" {
		oldHike.Description = updatedHike.Description
	}
	*updatedHike = oldHike
	result := r.db.Save(updatedHike)
	return result.Error
}
