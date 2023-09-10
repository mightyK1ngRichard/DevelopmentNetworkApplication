package repository

import (
	"VikingsServer/internal/app/ds"
)

func (r *Repository) CitiesList() (*[]ds.City, error) {
	rows, err := r.db.Query(`SELECT * FROM cities`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cities []ds.City
	for rows.Next() {
		c := ds.City{}
		if err := rows.Scan(
			&c.ID,
			&c.CityName,
			&c.Status,
			&c.ImageURL,
		); err != nil {
			r.logger.Error(err)
			continue
		}
		cities = append(cities, c)
	}

	return &cities, nil
}
