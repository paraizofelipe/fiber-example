package user

import (
	"github.com/jmoiron/sqlx"
)

type UserService struct {
	Repository Repository
}

func NewUserService(db *sqlx.DB) Service {
	return &UserService{
		Repository: NewUserRepo(db),
	}
}

func (r UserService) FindByEmail(email string) (user []User, err error) {
	if user, err = r.Repository.FindByEmail(email); err != nil {
		return
	}
	return
}

func (r UserService) FindById(id string) (user User, err error) {
	if user, err = r.Repository.FindById(id); err != nil {
		return
	}
	return
}

func (r UserService) Add(user User) (err error) {
	if err = r.Repository.Add(user); err != nil {
		return
	}
	return
}

func (r UserService) Update(user User) (err error) {
	if err = r.Repository.Update(user); err != nil {
		return
	}
	return
}

func (r UserService) DeleteById(id string) (err error) {
	if err = r.Repository.DeleteById(id); err != nil {
		return
	}
	return
}
