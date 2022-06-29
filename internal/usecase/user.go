package usecase

import "log"

func NewUserUseCase(db UserRepository, meili UserMeiliRepository) UserUseCase {
	return &_userUseCase{db: db, meili: meili}
}

func (u *_userUseCase) UpdateUser(id int) error {
	log.Println("requested updating user", id)
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
