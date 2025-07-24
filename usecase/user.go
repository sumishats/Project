package usecase

import (
	"sample/jwt"
	"sample/model"

	"gorm.io/gorm"
)

type UserUseCase struct {
	db *gorm.DB
}

func NewUserUseCase(db *gorm.DB) *UserUseCase {
	return &UserUseCase{db: db}

}

func (uc *UserUseCase) GetDb() *gorm.DB {
	return uc.db
}

func (uc *UserUseCase) Signup(user model.User) error {
	user.Role = "user"
	return uc.db.Create(&user).Error
}

func (uc *UserUseCase) Login(email, password string) (*model.User, string, error) {
	var user model.User
	err := uc.db.Where("email = ? AND password = ? AND role = ?", email, password, "user").First(&user).Error
	if err != nil {
		return nil, "", err
	}

	// Generate JWT token
	token, err := jwt.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}
