package product

import (
	"github.com/jmoiron/sqlx"
)

type ProductService struct {
	Repository Repository
}

func NewProductService(db *sqlx.DB) Service {
	return &ProductService{
		Repository: NewProductRepo(db),
	}
}

func (r ProductService) FindByTitle(title string) (product []Product, err error) {
	if product, err = r.Repository.FindByTitle(title); err != nil {
		return
	}
	return
}

func (r ProductService) FindById(id string) (product Product, err error) {
	if product, err = r.Repository.FindById(id); err != nil {
		return
	}
	return
}

func (r ProductService) Add(product Product) (err error) {
	if err = r.Repository.Add(product); err != nil {
		return
	}
	return
}

func (r ProductService) Update(product Product) (err error) {
	if err = r.Repository.Update(product); err != nil {
		return
	}
	return
}

func (r ProductService) DeleteById(id string) (err error) {
	if err = r.Repository.DeleteById(id); err != nil {
		return
	}
	return
}
