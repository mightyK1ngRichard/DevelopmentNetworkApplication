package repository

import (
	"VikingsServer/internal/app/ds"
)

func (r *Repository) AddViking(v *ds.Viking) error {
	return nil
}

func (r *Repository) UpdateViking(v *ds.Viking) error {
	return nil
}

// MARK: - Private filter sql command

/*
func addVikingSQLCommand(v *ds.Vikings) string {
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

	if v.CityOfBirthID != -1 {
		params += ", cityofbirth"
		values += fmt.Sprintf(", %d", v.CityOfBirthID)
	}

	if v.ImageURL != "" {
		params += ", imageurl"
		values += fmt.Sprintf(", '%s'", v.ImageURL)
	}

	// Поделил на части, чтоб псевдо ошибку не выдавало.
	command := fmt.Sprintf("INTO vikings (%s) VALUES (%s) RETURNING id;", params, values)
	return fmt.Sprintf("INSERT %s", command)
}

func updateVikingSQLCommand(v *ds.Vikings) string {
	params := fmt.Sprintf("id=%d", v.ID)
	if v.VikingName != "" {
		params += fmt.Sprintf(",vikingname='%s'", v.VikingName)
	}
	if v.Post != "" {
		params += fmt.Sprintf(",post='%s'", v.Post)
	}
	if v.Birthday != "" {
		params += fmt.Sprintf(",birthday='%s'", v.Birthday)
	}
	if v.DayOfDeath != "" {
		params += fmt.Sprintf(",dayofdeath='%s'", v.DayOfDeath)
	}
	if v.CityOfBirthID != -1 {
		params += fmt.Sprintf(",cityofbirthid=%d", v.CityOfBirthID)
	}
	if v.ImageURL != "" {
		params += fmt.Sprintf(",imageurl=%s", v.ImageURL)
	}
	updateString := `UPDATE vikings `
	setString := fmt.Sprintf(`SET %s WHERE id=%d;`, params, v.ID)
	return updateString + setString
}
*/
