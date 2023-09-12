package repository

import (
	"VikingsServer/internal/app/ds"
	"fmt"
)

func (r *Repository) AddViking(v *ds.Vikings) error {
	params := `vikingname, post`
	values := fmt.Sprintf("'%s', '%s'", v.VikingName, v.Post)

	if v.Birthday != "" {
		params += ", birthday"
		values += fmt.Sprintf(", '%s'", v.Birthday)
	}

	if v.DayOfDeath != "" {
		params += ", dayofdeath"
		values += fmt.Sprintf(", '%s'", v.DayOfDeath)
	}

	if v.CityOfBirth != -1 {
		params += ", cityofbirth"
		values += fmt.Sprintf(", %d", v.CityOfBirth)
	}

	if v.ImageURL != "" {
		params += ", imageurl"
		values += fmt.Sprintf(", '%s'", v.ImageURL)
	}

	// Поделил на части, чтоб ошибку не давало.
	command := fmt.Sprintf("INTO vikings (%s) VALUES (%s) RETURNING id;", params, values)
	sqlCommand := fmt.Sprintf("INSERT %s", command)

	row := r.db.QueryRow(sqlCommand)
	err := row.Scan(&v.ID)
	if err != nil {
		return err
	}
	return nil
}
