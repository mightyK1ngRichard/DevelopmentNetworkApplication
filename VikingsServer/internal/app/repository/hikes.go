package repository

import (
	"VikingsServer/internal/app/ds"
)

func (r *Repository) HikesList() (*[]ds.Hike, error) {
	var hikes []ds.Hike
	result := r.db.Preload("Author").Preload("Status").Find(&hikes)
	return &hikes, result.Error
}

func (r *Repository) AddHike(hike *ds.Hike) error {
	result := r.db.Create(&hike)
	return result.Error
}

func (r *Repository) DeleteHike(id uint) error {
	var hike ds.Hike
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
	var oldHike ds.Hike
	if result := r.db.First(&oldHike, updatedHike.ID); result.Error != nil {
		return result.Error
	}
	if updatedHike.HikeName != "" {
		oldHike.HikeName = updatedHike.HikeName
	}
	if updatedHike.DateEnd.String() != "0001-01-01 00:00:00 +0000 UTC" {
		oldHike.DateStart = updatedHike.DateStart
	}
	if updatedHike.DateEnd.String() != "0001-01-01 00:00:00 +0000 UTC" {
		oldHike.DateEnd = updatedHike.DateEnd
	}
	if updatedHike.DateStartPreparing.String() != "0001-01-01 00:00:00 +0000 UTC" {
		oldHike.DateStartPreparing = updatedHike.DateStartPreparing
	}
	if updatedHike.AuthorID != 0 {
		oldHike.AuthorID = updatedHike.AuthorID
	}
	if updatedHike.StatusID != 0 {
		oldHike.StatusID = updatedHike.StatusID
	}
	if updatedHike.Description != "" {
		oldHike.Description = updatedHike.Description
	}
	*updatedHike = oldHike
	result := r.db.Save(updatedHike)
	return result.Error
}
