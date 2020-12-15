package user

import (
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type UserRepo struct {
	db   *sqlx.DB
	psql sq.StatementBuilderType
}

func NewUserRepo(db *sqlx.DB) Repository {
	return &UserRepo{
		db:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (r UserRepo) FindByEmail(email string) (user User, err error) {
	query := r.psql.Select("*").
		From("users").
		Where(sq.Eq{"email": email})

	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := r.db.Preparex(statement)
	if err != nil {
		return
	}
	if err = stmt.Select(&user, args...); err != nil {
		return
	}
	err = stmt.Close()
	return
}

func (r UserRepo) FindById(id string) (user User, err error) {
	var users []User
	query := r.psql.Select("*").
		From("users").
		Where(sq.Eq{"id": id}).
		Limit(1)

	log.Println(query.ToSql())
	statement, args, err := query.ToSql()
	if err != nil {
		return
	}
	stmt, err := r.db.Preparex(statement)
	if err != nil {
		return
	}
	if err = stmt.Select(&users, args...); err != nil {
		return
	}
	err = stmt.Close()

	if len(users) != 0 {
		user = users[0]
	}

	return
}

func (r UserRepo) Add(user User) (err error) {
	query := r.psql.Insert("users").
		Columns("id", "username", "password", "name", "email", "roles").
		Values(
			uuid.NewV4().String(),
			user.Username,
			user.Password,
			user.Name,
			user.Email,
			user.Roles)

	query = query.RunWith(r.db)
	if _, err = query.Exec(); err != nil {
		log.Println(err)
		return
	}
	return
}

func (r UserRepo) Update(user User) (err error) {
	query := r.psql.Update("users").
		SetMap(sq.Eq{
			"email":    user.Email,
			"username": user.Username,
			"password": user.Password,
			"name":     user.Name,
			"roles":    user.Roles,
		}).
		Where(sq.Eq{
			"id": user.ID,
		})

	query = query.RunWith(r.db)
	log.Println(query.ToSql())
	if _, err = query.Exec(); err != nil {
		log.Println(err)
		return
	}
	return
}

func (r UserRepo) DeleteById(id string) (err error) {
	query := r.psql.Delete("users").
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
