package redis

import (
	"keeper/internal/usecase"
	"keeper/pkg/redispubhandler"
	"log"
	"strconv"
)

type beatmap_handler struct {
	t usecase.BeatmapsUseCase
}

// type result struct {
// 	UserID int `json:"user_id"`
// }

func NewBeatmapHandler(t usecase.BeatmapsUseCase) *beatmap_handler {
	return &beatmap_handler{
		t: t,
	}
}

func (b *beatmap_handler) Response(r *redispubhandler.Context) {
	if r.Error != nil {
		log.Fatal(r.Error)
	}
	beatmapID, err := strconv.Atoi(r.Message)
	if err == nil && beatmapID != 0 {
		err := b.t.UpdateBeatmapSet(beatmapID)
		if err != nil {
			log.Println(err)
		}
	}

}
