package repository

import (
	"backend-bangkit/entity"
	"backend-bangkit/pkg/errs"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(*entity.User) (*entity.User, errs.MessageErr)
	GetUserByEmail(email string) (*entity.User, errs.MessageErr)
}

type userPg struct {
	db *gorm.DB
}

func NewUserPg(db *gorm.DB) AuthRepository {
	return &userPg{db}
}

func (u *userPg) Register(user *entity.User) (*entity.User, errs.MessageErr) {

	if err := u.db.Create(user).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewInternalServerError("Failed to register new user")
	}

	return user, nil
}

func (u *userPg) GetUserByEmail(email string) (*entity.User, errs.MessageErr) {
	var user entity.User

	if err := u.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with email %s is not found", email))
	}

	return &user, nil
}
