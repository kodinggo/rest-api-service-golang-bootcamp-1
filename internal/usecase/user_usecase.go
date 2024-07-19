package usecase

import (
	"kodinggo/internal/model"

	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	userRepo model.IUserRepository
	log      *logrus.Logger
}

func NewUserUsecase(
	userRepo model.IUserRepository,
	log *logrus.Logger,
) model.IUserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Create(user model.User) error {
	log := u.log.WithFields(logrus.Fields{
		"user": user,
	})

	err := u.userRepo.Create(user)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (u *UserUsecase) Login(username string, password string) (model.User, error) {
	log := u.log.WithFields(logrus.Fields{
		"username": username,
		"password": password,
	})

	user, err := u.userRepo.Login(username)
	if err != nil {
		log.Error(err)
		return model.User{}, model.ErrUsernameNotFound
	}

	if !user.IsPasswordMatch(password) {
		log.Error(err)
		return model.User{}, model.ErrInvalidPassword
	}

	return user, nil
}

func (u *UserUsecase) FindByUsername(username string) (model.User, error) {
	log := u.log.WithFields(logrus.Fields{
		"username": username,
	})

	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		log.Error(err)
		return model.User{}, err
	}

	return user, nil
}
