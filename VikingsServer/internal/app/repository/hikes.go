package repository

import (
	"VikingsServer/internal/app/ds"
)

func (r *Repository) HikesList() (*[]ds.Hikes, error) {
	rows, err := r.db.Query(`SELECT * FROM hikes`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var hikes []ds.Hikes
	for rows.Next() {
		h := ds.Hikes{}
		if err := rows.Scan(
			&h.ID,
			&h.HikeName,
			&h.DateStart,
			&h.DateEnd,
			&h.AuthorID,
			&h.Status,
			&h.Description,
		); err != nil {
			r.logger.Error(err)
			continue
		}
		hikes = append(hikes, h)
	}

	return &hikes, nil
}
