package http

import (
	"fmt"
	"log"
	"net/http"
	"test_backend_1/domain/model"
	"test_backend_1/usecase"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	Login(c *gin.Context)
}

type UserHandler struct {
	userUsecase usecase.IUserUsecase
}

func NewUserHandler(userUsecase usecase.IUserUsecase) IUserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (userHandler *UserHandler) Login(c *gin.Context) {
	var req model.ReqLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("An error occurred: %v", err)
		c.JSON(http.StatusBadRequest, fmt.Sprintf("An error occurred: %v", err.Error()))
	}

	res := userHandler.userUsecase.Login(c.Request.Context(), req)

	c.JSON(http.StatusOK, res)
}
