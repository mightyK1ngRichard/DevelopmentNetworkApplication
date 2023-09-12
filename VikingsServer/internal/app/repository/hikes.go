package repository

import (
	"VikingsServer/internal/app/ds"
)

func (r *Repository) HikesList() (*[]ds.Hikes, error) {
	sqlCommand := ` SELECT h.id, h.hikename, h.datestart, h.dateend, v.vikingname, hs.name, h.description
					FROM hikes h
					LEFT JOIN hikestatuses hs ON h.status = hs.id
					LEFT JOIN vikings v ON v.id = h.leader;`

	rows, err := r.db.Query(sqlCommand)
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
			&h.Leader,
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
