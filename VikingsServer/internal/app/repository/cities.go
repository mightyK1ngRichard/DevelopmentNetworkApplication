package repository

import (
	"VikingsServer/internal/app/ds"
	"database/sql"
)

func (r *Repository) CitiesList() (*[]ds.City, error) {
	sqlCommand := `SELECT c.id, c.cityname, s.name, c.imageurl, c.description FROM cities c LEFT JOIN citystatuses s ON c.status = s.id;`
	rows, err := r.db.Query(sqlCommand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cities []ds.City
	for rows.Next() {
		c := ds.City{}
		var description sql.NullString
		if err := rows.Scan(
			&c.ID,
			&c.CityName,
			&c.Status,
			&c.ImageURL,
			&description,
		); err != nil {
			r.logger.Error(err)
			continue
		}
		if description.Valid {
			c.Description = description.String
		} else {
			c.Description = "Описание отсутствует"
		}
		cities = append(cities, c)
	}

	return &cities, nil
}
