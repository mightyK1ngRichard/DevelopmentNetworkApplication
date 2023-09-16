package repository

import (
	"VikingsServer/internal/app/ds"
)

func (r *Repository) CitiesList() (*[]ds.City, error) {
	var cities []ds.City
	r.db.Preload("Status").Where(
		`status_id IN (SELECT id FROM city_statuses WHERE status_name = ? LIMIT 1)`,
		"существует",
	).Find(&cities)
	return &cities, nil
}

func (r *Repository) DeleteCity(id uint) error {
	var newStatus ds.CityStatus
	r.db.Where("status_name = ?", "уничтожен").First(&newStatus)

	if newStatus.ID != 0 {
		var city ds.City
		r.db.First(&city, id)
		city.StatusID = newStatus.ID
		r.db.Save(&city)
	}
	return nil
}

/*
func (r *Repository) DeleteCity(id int) error {
	sqlCommand := `DELETE FROM cities WHERE id=$1;`
	_, err := r.db.Exec(sqlCommand, id)
	if err != nil {
		r.logger.Error(err)
		return err
	}
	return nil
}

func (r *Repository) DeleteCityWithStatus(id int) error {
	sqlCommand := `UPDATE cities SET status=2 WHERE id = $1;`
	_, err := r.db.Exec(sqlCommand, id)
	if err != nil {
		return err
	}
	return nil
}
*/
