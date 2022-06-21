package user_db

import (
	"keeper/internal/entity"
	rep "keeper/internal/repository"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) rep.UserRepository {
	return &repository{db: db}
}

func (r *repository) GetUsers(name string) ([]entity.User, error) {
	rows, err := r.db.Queryx("SELECT id, users.username, country, privileges, beta_key, email, username_aka FROM users LEFT JOIN users_stats USING (id) WHERE users.username LIKE ?", name)
	if err != nil {
		return nil, err
	}
	reslut := []entity.User{}
	for rows.Next() {
		u := entity.User{}
		rows.StructScan(&u)
		reslut = append(reslut, u)
	}
	return reslut, nil
}

func (r *repository) GetUserByID(id int) (*entity.User, error) {
	row := r.db.QueryRowx("SELECT id, users.username, country, privileges, beta_key, email, username_aka FROM users LEFT JOIN users_stats USING (id) WHERE users.id = ?", id)
	result := entity.User{}
	row.StructScan(&result)
	return &result, nil
}
