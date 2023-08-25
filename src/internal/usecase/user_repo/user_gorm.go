package user_repo

import (
	"gorm.io/gorm"
	"users-segments-service/internal/entity"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) Create(user *entity.User) (*entity.User, error) {
	return user, u.db.Create(&user).Error
}

func (u *UserRepository) UserExists(email string) (bool, error) {
	var count int64
	err := u.db.Model(&entity.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (u *UserRepository) UserExistsByID(userID uint) (bool, error) {
	var count int64
	err := u.db.Model(&entity.User{}).Where("id = ?", userID).Count(&count).Error
	return count > 0, err
}
