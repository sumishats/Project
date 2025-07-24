package helper

import (
	"sample/model"

	"gorm.io/gorm"
)

func EmailExist(db *gorm.DB,email string)(bool,error){
	var user model.User
	//select * from users where email =''
	err:=db.Where("email=?",email).First(&user).Error
	if err!=nil{
		return false,err
	}
	return true,nil
}