package beatmap_db

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

func (r *repository) GetBeatmapByID(id int) (*entity.BeatmapSet, error) {
	row := r.db.QueryRowx("SELECT id, users.username, country, privileges, beta_key, email, username_aka FROM users LEFT JOIN users_stats USING (id) WHERE users.id = ?", id)
	result := entity.User{}
	row.StructScan(&result)
	return &result, nil
}
