package usecase

import (
	"fmt"
	"gorm.io/gorm"
	"users-segments-service/internal/entity"
)

type UserUseCase struct {
	userRepo UserRepo
}

func NewUserUseCase(ur UserRepo) *UserUseCase {
	return &UserUseCase{
		userRepo: ur,
	}
}

func (u UserUseCase) Register(email string, name string, lastName string) (*entity.User, error) {
	exists, err := u.userRepo.UserExists(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("UserUsecase - Register - u.userRepo.UserExists")
	}
	if exists {
		return nil, ErrorUserAlreadyExists
	}

	return u.userRepo.Create(&entity.User{
		Email:    email,
		Name:     name,
		LastName: lastName,
	})
}
