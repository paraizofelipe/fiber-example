package user

import "encoding/json"

// User ---
type User struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Name     string `json:"name" db:"name"`
	Email    string `json:"email" db:"email"`
	Roles    string `json:"roles" db:"roles"`
}

type Reader interface {
	FindByEmail(email string) (User, error)
	FindById(id string) (User, error)
}

type Writer interface {
	Add(user User) error
	Update(user User) error
	DeleteById(id string) error
}

type Repository interface {
	Reader
	Writer
}

type Service interface {
	Reader
	Writer
}

func (m User) String() string {
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(b)
}
