package handler

import (
	"net/http"
	"sample/model"
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
//admin login
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
// Get Users handles get all users request
func (h *AdminHandler) GetUsers(c *gin.Context) {
	users, err := h.AdminUseCase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
// CreateUser handles create user request
func (h *AdminHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AdminUseCase.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
		"user":    user,
	})
}

// GetUser handles get single user request
func (h *AdminHandler) GetUser(c *gin.Context) {
	id := c.Param("id")

	user, err := h.AdminUseCase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// UpdateUser handles update user request
func (h *AdminHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.AdminUseCase.UpdateUser(id, &input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated",
		"user":    user,
	})
}

// DeleteUser handles delete user request
func (h *AdminHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := h.AdminUseCase.DeleteUser(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}

func(h *AdminHandler)Home(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"welcome to home page"})
}