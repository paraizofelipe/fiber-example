package seller

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type SellerRepo struct {
	db   *sqlx.DB
	psql sq.StatementBuilderType
}

func NewSellerRepo(db *sqlx.DB) Repository {
	return &SellerRepo{
		db:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r SellerRepo) FindByName(title string) (seller Seller, err error) {
	return
}

func (r SellerRepo) FindById(id string) (seller Seller, err error) {
	return
}

func (r SellerRepo) Add(seller Seller) (err error) {
	return
}

func (r SellerRepo) Update(seller Seller) (err error) {
	return
}

func (r SellerRepo) DeleteById(id string) (err error) {
	return
}
