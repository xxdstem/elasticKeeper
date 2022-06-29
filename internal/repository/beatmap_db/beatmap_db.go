package beatmap_db

import (
	"keeper/internal/entity"
	rep "keeper/internal/repository"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) rep.BeatmapRepository {
	return &repository{db: db}
}

func (r *repository) GetBeatmapByID(id int) (*entity.BeatmapSet, error) {
	rows := r.db.QueryRowx(`SELECT favourite_count, artist, title, creator, beatmapset_id,
    IF(playcount is NULL, 0, SUM(playcount)) playcount, IF(passcount is NULL, 0, SUM(passcount)) passcount,
    ranked, ranked_status_freezed, ranking_data FROM beatmaps
    LEFT JOIN beatmap_plays USING (beatmap_md5)
    WHERE beatmapset_id = ? GROUP BY beatmapset_id ORDER BY ranking_data DESC`, id)
	var beatmap entity.BeatmapSet
	rows.StructScan(&beatmap)
	diffs, err := r.db.Queryx("SELECT ar, od, hp, cs, bpm, IF(total_length > 0, total_length, hit_length) total_length, mode, beatmap_id, version, difficulty_std, difficulty_mania, difficulty_taiko, difficulty_ctb FROM beatmaps WHERE beatmapset_id = ? ORDER BY difficulty_std ASC, difficulty_taiko ASC, difficulty_ctb ASC, difficulty_mania ASC", beatmap.ID)
	if err != nil {
		return nil, err
	}
	for diffs.Next() {
		var diff entity.BeatmapDiff
		diffs.StructScan(&diff)
		beatmap.Beatmaps = append(beatmap.Beatmaps, diff)
	}

	return &beatmap, nil
}
