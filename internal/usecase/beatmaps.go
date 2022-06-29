package usecase

import "log"

func NewBeatmapsUseCase(db BeatmapRepository, meili BeatmapMeiliRepository) BeatmapsUseCase {
	return &_beatmapUseCase{db: db, meili: meili}
}

func (u *_beatmapUseCase) UpdateBeatmapSet(id int) error {
	log.Println("requested updating beatmapset ", id)
	b, err := u.db.GetBeatmapByID(id)
	if err != nil {
		return err
	}
	u.meili.UpdateBeatmap(b)
	return nil
}
