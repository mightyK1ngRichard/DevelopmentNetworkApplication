package repository

import (
	"VikingsServer/internal/app/ds"
)

func (r *Repository) HikesList() (*[]ds.Hike, error) {
	var hikes []ds.Hike
	r.db.Preload("Author").Preload("Status").Find(&hikes)
	return &hikes, nil
}

/*
func (r *Repository) HikesList() (*[]ds.Hike, error) {
	rows, err := r.db.Query(`SELECT * FROM hikes`)
>>>>>>> backend/lab02
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var hikes []ds.Hike
	for rows.Next() {
		h := ds.Hike{}
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
*/
