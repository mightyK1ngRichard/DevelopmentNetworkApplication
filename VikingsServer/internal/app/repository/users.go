package repository

import "VikingsServer/internal/app/ds"

func (r *Repository) UsersList() (*[]ds.User, error) {
	var users []ds.User
	result := r.db.Find(&users)
	return &users, result.Error
}
