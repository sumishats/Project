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