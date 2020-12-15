package product

import "encoding/json"

// Product ---
type Product struct {
	ID          string `json:"id,omitempty" validate:"required" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" validate:"required" db:"description"`
	Amount      int    `json:"amount" validate:"required,number" db:"amount"`
}

type Reader interface {
	FindByTitle(title string) ([]Product, error)
	FindById(id string) (Product, error)
}

type Writer interface {
	Add(product Product) error
	Update(product Product) error
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

func (m Product) String() string {
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(b)
}
