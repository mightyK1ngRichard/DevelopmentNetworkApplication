package repository

import "VikingsServer/internal/app/ds"

func (r *Repository) AuthorsList() (*[]ds.Author, error) {
	var authors []ds.Author
	result := r.db.Find(&authors)
	return &authors, result.Error
}
