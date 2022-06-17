package entity

type BeatmapSet struct {
	ID             int `json:"beatmapset_id"`
	FavouriteCount int `json:"favourite_count"`
	Artist         string
	Title          string
	Creator        string
	PlayCount      int
	PassCoutn      int
	Ranked         int8
	RankFreezed    int8 `json:"ranked_status_freezed"`
	RankDate       int  `json:"ranking_data"`
	Beatmaps       []struct {
		ID              int `json:"beatmap_id"`
		AR              int
		OD              int
		HP              int
		CS              int
		BPM             int
		Length          int `json:"total_length"`
		Mode            int
		Version         string
		DifficultyStd   float64 `json:"difficulty_std"`
		DifficultyTaiko float64 `json:"difficulty_taiko"`
		DifficultyCtb   float64 `json:"difficulty_ctb"`
		DifficultyMania float64 `json:"difficulty_mania"`
	}
}
