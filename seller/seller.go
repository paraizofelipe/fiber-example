package seller

import (
	"encoding/json"

	"github.com/paraizofelipe/fiber-example/product"
)

// Seller ---
type Seller struct {
	Name     string            `json:"name"`
	Address  string            `json:"address"`
	Products []product.Product `json:"products"`
}

type Reader interface {
	FindByName(email string) (Seller, error)
	FindById(id string) (Seller, error)
}

type Writer interface {
	Add(seller Seller) error
	Update(seller Seller) error
	DeleteById(id string) error
}

type Repository interface {
	Reader
	Writer
}

func (m Seller) String() string {
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(b)
}
