package usecase

import (
	"sample/jwt"
	"sample/model"

	"gorm.io/gorm"
)

type AdminUseCase struct{
	db *gorm.DB
}

func NewAdminUseCase(db *gorm.DB)*AdminUseCase{
	return &AdminUseCase{db: db}
}
func (uc *AdminUseCase)AdminLogin(email string,password string)(*model.User,string,error){
	var user model.User
	err:=uc.db.Where("email = ?  And  password = ? And  role = ?",email,password,"admin").First(&user).Error
	if err!=nil{
		return &model.User{},"",err
	}
	token,err:=jwt.GenerateToken(user.ID,email,user.Role)
	if err!=nil{
		return &model.User{},"",err
	}
	return &user,token,nil
	
}
// GetAllUsers gets all users
func (uc *AdminUseCase) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := uc.db.Find(&users).Error
	return users, err
}

// CreateUser creates a new user
func (uc *AdminUseCase) CreateUser(user *model.User) error {
	return uc.db.Create(user).Error
}

// GetUserByID gets a user by ID
func (uc *AdminUseCase) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := uc.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user
func (uc *AdminUseCase) UpdateUser(id string, input *model.User) (*model.User, error) {
	var user model.User
	if err := uc.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	uc.db.Model(&user).Updates(input)
	return &user, nil
}

// DeleteUser deletes a user
func (uc *AdminUseCase) DeleteUser(id string) error {
	// First check if user exists
	var user model.User
	if err := uc.db.First(&user, id).Error; err != nil {
		return err // User not found
	}

	// Delete the user
	return uc.db.Delete(&user).Error
}