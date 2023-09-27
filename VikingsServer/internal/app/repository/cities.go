package repository

import (
	"VikingsServer/internal/app/ds"
	"fmt"
)

func (r *Repository) CitiesList() (*[]ds.City, error) {
	var cities []ds.City
	result := r.db.Preload("Status").Where(
		`status_id IN (SELECT id FROM city_statuses WHERE status_name = ? LIMIT 1)`,
		"существует",
	).Find(&cities)
	return &cities, result.Error
}

func (r *Repository) CitiesById(id uint) (*ds.City, error) {
	city := ds.City{}
	result := r.db.Preload("Status").First(&city, id)
	return &city, result.Error
}

func (r *Repository) DeleteCity(id uint) error {
	var newStatus ds.CityStatus
	if result := r.db.Where("status_name = ?", "уничтожен").First(&newStatus); result.Error != nil {
		return result.Error
	}
	if newStatus.ID == 0 {
		return fmt.Errorf("city status not found. may be it's name was changed by someone")
	}
	var city ds.City
	if result := r.db.First(&city, id); result.Error != nil {
		return result.Error
	}
	if city.ID == 0 {
		return fmt.Errorf("city not found")
	}
	city.StatusID = newStatus.ID
	result := r.db.Save(&city)
	return result.Error
}

func (r *Repository) AddCity(city *ds.City) error {
	result := r.db.Create(&city)
	return result.Error
}
