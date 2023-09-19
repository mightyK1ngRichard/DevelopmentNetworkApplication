package repository

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/utils"
)

func (r *Repository) HikesList() (*[]ds.Hike, error) {
	var hikes []ds.Hike
	result := r.db.Preload("Participants.Viking.CityOfBirth.Status").Preload("DestinationHikes.City.Status").Preload("Author").Preload("Status").Find(&hikes)
	return &hikes, result.Error
}

func (r *Repository) HikeByID(id uint) (*ds.Hike, error) {
	hike := ds.Hike{}
	result := r.db.Preload("Author").Preload("Status").First(&hike, id)
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
	if result := r.db.Where("status_name = ?", "удалён").First(&newStatus); result.Error != nil {
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
	if updatedHike.HikeName != utils.EmptyString {
		oldHike.HikeName = updatedHike.HikeName
	}
	if updatedHike.DateEnd.String() != utils.EmptyDate {
		oldHike.DateStart = updatedHike.DateStart
	}
	if updatedHike.DateEnd.String() != utils.EmptyDate {
		oldHike.DateEnd = updatedHike.DateEnd
	}
	if updatedHike.DateStartPreparing.String() != utils.EmptyDate {
		oldHike.DateStartPreparing = updatedHike.DateStartPreparing
	}
	if updatedHike.AuthorID != utils.EmptyInt {
		oldHike.AuthorID = updatedHike.AuthorID
	}
	if updatedHike.StatusID != utils.EmptyInt {
		oldHike.StatusID = updatedHike.StatusID
	}
	if updatedHike.Description != utils.EmptyString {
		oldHike.Description = updatedHike.Description
	}
	*updatedHike = oldHike
	result := r.db.Save(updatedHike)
	return result.Error
}
