package handler

import (
	"net/http"
	"sample/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type AdminHandler struct{
	AdminUseCase *usecase.AdminUseCase
}
func NewAdminHandler(db *gorm.DB)*AdminHandler{
	return &AdminHandler{AdminUseCase: usecase.NewAdminUseCase(db)}

}
func (h *AdminHandler)AdminLogin(c *gin.Context){
	var input struct{
		Email string  `json:"email"`
		Password string `json:"password"`
	}
	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return 
	}

	user,token,err:=h.AdminUseCase.AdminLogin(input.Email,input.Password)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"AdminLogin successfuly","user":user,"token":token})
	
}