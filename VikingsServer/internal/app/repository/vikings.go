package repository

import (
	"VikingsServer/internal/app/ds"
	"VikingsServer/internal/utils"
)

func (r *Repository) VikingList() (*[]ds.Viking, error) {
	var vikings []ds.Viking
	param := r.db.Preload("CityOfBirth").Preload("CityOfBirth.Status").Find(&vikings)
	return &vikings, param.Error
}

func (r *Repository) VikingById(id uint) (*ds.Viking, error) {
	viking := ds.Viking{}
	result := r.db.Preload("CityOfBirth.Status").First(&viking, id)
	return &viking, result.Error
}

func (r *Repository) AddViking(viking *ds.Viking) error {
	result := r.db.Create(viking)
	return result.Error
}

func (r *Repository) UpdateViking(updatedViking *ds.Viking) error {
	oldViking := ds.Viking{}
	if result := r.db.First(&oldViking, updatedViking.ID); result.Error != nil {
		return result.Error
	}
	if updatedViking.VikingName != utils.EmptyString {
		oldViking.VikingName = updatedViking.VikingName
	}
	if updatedViking.Post != utils.EmptyString {
		oldViking.Post = updatedViking.Post
	}
	if updatedViking.Birthday.String() != utils.EmptyDate {
		oldViking.Birthday = updatedViking.Birthday
	}
	if updatedViking.DayOfDeath.String() != utils.EmptyDate {
		oldViking.DayOfDeath = updatedViking.DayOfDeath
	}
	if updatedViking.CityOfBirthID != utils.EmptyInt {
		oldViking.CityOfBirthID = updatedViking.CityOfBirthID
	}
	if updatedViking.ImageURL != utils.EmptyString {
		oldViking.ImageURL = updatedViking.ImageURL
	}
	*updatedViking = oldViking
	result := r.db.Save(updatedViking)
	return result.Error
}

func (r *Repository) DeleteViking(id uint) error {
	return nil
}
