package repository

import "VikingsServer/internal/app/ds"

func (r *Repository) DestinationHikesList() (*[]ds.DestinationHikes, error) {
	var destinationHikes []ds.DestinationHikes
	result := r.db.Preload("Hike.Status").Preload("City.Status").Find(&destinationHikes)
	return &destinationHikes, result.Error
}

func (r *Repository) AddDestinationToHike(dh *ds.DestinationHikes) error {
	result := r.db.Create(dh)
	return result.Error
}
