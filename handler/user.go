package handler

import (
	"net/http"
	"sample/helper"
	"sample/jwt"
	"sample/model"
	"sample/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct{
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(db *gorm.DB)*UserHandler{
	return &UserHandler{userUseCase:usecase.NewUserUseCase(db)}
}

//User signup 
func (h *UserHandler)Signup(c *gin.Context){

	var user model.User

	if err:=c.ShouldBindJSON(&user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return 
	}

	emailexist,_:=helper.EmailExist(h.userUseCase.GetDb(),user.Email)

	if emailexist{
		c.JSON(http.StatusBadRequest,gin.H{"error":"email already exist"})
	}
	if err:=h.userUseCase.Signup(user);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	//gernarate token
	token,err:=jwt.GenerateToken(user.ID,user.Email,user.Role)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"failed to generate token"})
		return 
	}
	c.JSON(http.StatusOK,gin.H{"message":"signup successful","user":user,"token":token})
	
}
//user login 
func (h *UserHandler)Login(c *gin.Context){
	
	var userLogin struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err:=c.ShouldBindJSON(&userLogin);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return 
	}
	// Use usecase to handle business logic
	user, token, err := h.userUseCase.Login(userLogin.Email, userLogin.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
		"token":   token,
	})
}
	
