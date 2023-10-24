package repository

import (
	"VikingsServer/internal/app/ds"
)

func (r *Repository) UsersList() (*[]ds.User, error) {
	var users []ds.User
	result := r.db.Find(&users)
	return &users, result.Error
}

func (r *Repository) Register(user *ds.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) GetUserByLogin(login string) (*ds.User, error) {
	user := &ds.User{}

	if err := r.db.Where("login = ?", login).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
