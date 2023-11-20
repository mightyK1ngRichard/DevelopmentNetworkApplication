package repository

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/utils"
	"errors"
	"time"
)

func (r *Repository) HikesList(statusID string, startDate time.Time, endDate time.Time) (*[]ds.Hike, error) {
	var hikes []ds.Hike
	result := r.db.Preload("Status").
		Preload("DestinationHikes.City.Status").
		Preload("DestinationHikes.Hike.Status").
		Preload("User").
		Preload("Status").
		//Where("status_id = ?", statusID).Find(&hikes)
		Where("status_id = ? AND date_start_of_processing BETWEEN ? AND ?", statusID, startDate, endDate).
		Find(&hikes)
	return &hikes, result.Error
}

func (r *Repository) AddCityIntoHike(cityID uint, userID uint, serialNumber int) (uint, error) {
	hikeID, err := r.HikeBasketId()

	/// Корзины нет. Создадим заявку
	if err != nil {
		newHike := ds.Hike{
			UserID:      userID,
			DateCreated: time.Now(),
			StatusID:    1,
		}
		if errCreate := r.db.Create(&newHike).Error; errCreate != nil {
			return 0, errCreate
		}

		dh := ds.DestinationHikes{
			CityID:       cityID,
			HikeID:       newHike.ID,
			SerialNumber: serialNumber,
		}

		return dh.ID, r.AddDestinationToHike(&dh)
	}

	/// Корзина есть, добавляем туда
	dh := ds.DestinationHikes{
		CityID:       cityID,
		HikeID:       hikeID,
		SerialNumber: serialNumber,
	}

	return dh.ID, r.AddDestinationToHike(&dh)
}

func (r *Repository) HikeBasketId() (uint, error) {
	var hike ds.Hike
	result := r.db.Preload("Status").Where("status_id = ?", 1).First(&hike)
	return hike.ID, result.Error
}

func (r *Repository) HikeByID(id uint) (*ds.Hike, error) {
	hike := ds.Hike{}
	result := r.db.Preload("User").Preload("DestinationHikes.Hike.Status").Preload("DestinationHikes.Hike.User").Preload("DestinationHikes.City.Status").Preload("Status").First(&hike, id)
	return &hike, result.Error
}

func (r *Repository) AddHike(hike *ds.Hike) error {
	return r.db.Create(&hike).Error
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

func (r *Repository) UpdateStatusForUser(hikeID uint, newStatusID uint) error {
	updatedHike := ds.Hike{}
	if result := r.db.First(&updatedHike, hikeID); result.Error != nil {
		return result.Error
	}
	updatedHike.StatusID = newStatusID
	updatedHike.DateStartOfProcessing = time.Now()
	return r.db.Save(&updatedHike).Error
}

func (r *Repository) UpdateHikeForModerator(hikeID uint, newStatusID uint) error {
	updatedHike := ds.Hike{}
	if result := r.db.First(&updatedHike, hikeID); result.Error != nil {
		return result.Error
	}
	updatedHike.StatusID = newStatusID
	updatedHike.DateApprove = time.Now()
	return r.db.Save(&updatedHike).Error
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
		//oldHike.DateCreated = updatedHike.DateCreated
		return errors.New(`DateCreated cannot be updated`)
	}
	if updatedHike.DateStartOfProcessing.String() != utils.EmptyDate {
		//oldHike.DateStartOfProcessing = updatedHike.DateStartOfProcessing
		return errors.New(`DateStartOfProcessing cannot be updated`)
	}
	if updatedHike.DateApprove.String() != utils.EmptyDate {
		//oldHike.DateApprove = updatedHike.DateApprove
		return errors.New(`DateApprove cannot be updated`)
	}
	if updatedHike.DateEnd.String() != utils.EmptyDate {
		oldHike.DateEnd = updatedHike.DateEnd
	}
	if updatedHike.DateStartHike.String() != utils.EmptyDate {
		oldHike.DateStartHike = updatedHike.DateStartHike
	}
	if updatedHike.UserID != utils.EmptyInt {
		oldHike.UserID = updatedHike.UserID
	}
	if updatedHike.Description != "" {
		oldHike.Description = updatedHike.Description
	}
	//if updatedHike.StatusID != utils.EmptyInt {
	//	oldHike.StatusID = updatedHike.StatusID
	//}

	/// Меняем статус заявки на ожидания апрува модератора после редактирования
	oldHike.StatusID = 2
	*updatedHike = oldHike
	result := r.db.Save(updatedHike)
	return result.Error
}
