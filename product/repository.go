package product

import (
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type ProductRepo struct {
	db   *sqlx.DB
	psql sq.StatementBuilderType
}

func NewProductRepo(db *sqlx.DB) Repository {
	return &ProductRepo{
		db:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r ProductRepo) FindByTitle(title string) (products []Product, err error) {
	query := r.psql.Select("*").
		From("product").
		Where(sq.Eq{"title": title})

	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := r.db.Preparex(statement)
	if err != nil {
		return
	}
	if err = stmt.Select(&products, args...); err != nil {
		return
	}
	err = stmt.Close()
	return
}

func (r ProductRepo) FindById(id string) (product Product, err error) {
	var products []Product
	query := r.psql.Select("*").
		From("product").
		Where(sq.Eq{"id": id}).
		Limit(1)

	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := r.db.Preparex(statement)
	if err != nil {
		return
	}
	if err = stmt.Select(&products, args...); err != nil {
		return
	}
	err = stmt.Close()

	if len(products) != 0 {
		product = products[0]
	}

	return
}

func (r ProductRepo) Add(product Product) (err error) {
	query := r.psql.Insert("product").
		Columns("id", "title", "description", "amount").
		Values(
			uuid.NewV4().String(),
			product.Title,
			product.Description,
			product.Amount)

	query = query.RunWith(r.db)
	if _, err = query.Exec(); err != nil {
		log.Println(err)
		return
	}
	return
}

func (r ProductRepo) Update(product Product) (err error) {
	query := r.psql.Update("product").
		SetMap(sq.Eq{
			"title":       product.Title,
			"description": product.Description,
			"amount":      product.Amount,
		}).
		Where(sq.Eq{
			"id": product.ID,
		})

	query = query.RunWith(r.db)
	log.Println(query.ToSql())
	if _, err = query.Exec(); err != nil {
		log.Println(err)
		return
	}
	return
}

func (r ProductRepo) DeleteById(id string) (err error) {
	query := r.psql.Delete("product").
		Where(sq.Eq{
			"id": id,
		})

	query = query.RunWith(r.db)
	if _, err = query.Exec(); err != nil {
		log.Println(err)
		return
	}
	return
}
