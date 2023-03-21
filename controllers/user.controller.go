package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karncomsci/kasetwisai-shop/models"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.Users)

	userResponse := &models.UserResponse{
		ID:        currentUser.ID,
		Firstname: currentUser.Firstname,
		Lastname:  currentUser.Lastname,
		Email:     currentUser.Email,
		Photo:     currentUser.Photo,
		Role:      currentUser.Role,
		Provider:  currentUser.Provider,
		CreatedAt: currentUser.CreatedAt,
		UpdatedAt: currentUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

func (pc *UserController) CreateUser(ctx *gin.Context) {
	//currentUser := ctx.MustGet("currentUser").(models.Users)
	var payload *models.Users

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newUser := models.Users{
		ID:        payload.ID,
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
		Username:  payload.Username,
		Password:  payload.Password,
		Email:     payload.Email,
		Photo:     payload.Photo,
		Role:      payload.Role,
		Provider:  payload.Provider,
		Verified:  payload.Verified,
		Status:    payload.Status,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := pc.DB.Create(&newUser)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newUser})
}
