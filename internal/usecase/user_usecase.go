package usecase

import (
	"kodinggo/internal/model"
)

type UserUsecase struct {
	userRepo model.IUserRepository
}

func NewUserUsecase(userRepo model.IUserRepository) model.IUserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Create(user model.User) error {
	return u.userRepo.Create(user)
}

func (u *UserUsecase) Login(username string, password string) (model.User, error) {
	user, err := u.userRepo.Login(username)
	if err != nil {
		return model.User{}, model.ErrUsernameNotFound
	}

	if !user.IsPasswordMatch(password) {
		return model.User{}, model.ErrInvalidPassword
	}

	return user, nil
}

func (u *UserUsecase) FindByUsername(username string) (model.User, error) {
	return u.userRepo.FindByUsername(username)
}
