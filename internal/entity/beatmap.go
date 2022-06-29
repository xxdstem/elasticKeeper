package entity

type BeatmapSet struct {
	ID             int           `json:"beatmapset_id" db:"beatmapset_id"`
	FavouriteCount int           `json:"favourite_count" db:"favourite_count"`
	Artist         string        `json:"artist" db:"artist"`
	Title          string        `json:"title" db:"title"`
	Creator        string        `json:"creator" db:"creator"`
	PlayCount      int           `json:"playcount" db:"playcount"`
	PassCount      int           `json:"passcount" db:"passcount"`
	Ranked         int8          `json:"ranked" db:"ranked"`
	RankFreezed    int8          `json:"ranked_status_freezed" db:"ranked_status_freezed"`
	RankDate       int           `json:"ranking_data" db:"ranking_data"`
	Beatmaps       []BeatmapDiff `json:"beatmaps"`
}

type BeatmapDiff struct {
	ID              int     `json:"beatmap_id" db:"beatmap_id"`
	AR              float64 `json:"ar" db:"ar"`
	OD              float64 `json:"od" db:"od"`
	HP              float64 `json:"hp" db:"hp"`
	CS              float64 `json:"cs" db:"cs"`
	BPM             int     `json:"bpm" db:"bpm"`
	Length          int     `json:"total_length" db:"total_length"`
	Mode            int     `json:"mode" db:"mode"`
	Version         string  `json:"version" db:"version"`
	DifficultyStd   float64 `json:"difficulty_std" db:"difficulty_std"`
	DifficultyTaiko float64 `json:"difficulty_taiko" db:"difficulty_taiko"`
	DifficultyCtb   float64 `json:"difficulty_ctb" db:"difficulty_ctb"`
	DifficultyMania float64 `json:"difficulty_mania" db:"difficulty_mania"`
}
