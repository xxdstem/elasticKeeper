package usecase

import "log"

func NewBeatmapsUseCase(db UserRepository, meili UserMeiliRepository) BeatmapsUseCase {
	return &_useCase{db: db, meili: meili}
}

func (u *_useCase) UpdateBeatmapSet(id int) error {
	log.Println("requested updating beatmap", id)
	user, err := u.db.GetUserByID(id)
	if err != nil {
		log.Println(err)
		return err
	}
	err = u.meili.UpdateUser(user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
