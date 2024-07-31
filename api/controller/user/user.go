package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"test-backend-developer-sagala/application/user"
	responseMessage "test-backend-developer-sagala/common/response-message"
	authenticationDto "test-backend-developer-sagala/contract/authentication-dto"
)

type userController struct {
	userService user.UserService
}

func NewUserController(userService user.UserService) *userController {
	return &userController{userService}
}

func (h *userController) Register(c *gin.Context) {
	var input authenticationDto.Register
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := responseMessage.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := responseMessage.GetResponse(http.StatusBadRequest, false, errorMessage, false)
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := h.userService.Register(input)
	c.JSON(result.StatusCode, result)
}

func (h *userController) Login(c *gin.Context) {
	var input authenticationDto.Login
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errorMessages := responseMessage.BindRequestErrorChecking(err)

		errorMessage := strings.Join(errorMessages, ";")
		res := responseMessage.GetResponse(http.StatusUnauthorized, false, errorMessage, false)
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	result := h.userService.Login(input)
	c.JSON(result.StatusCode, result)
}
