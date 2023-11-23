package handler

import (
	"net/http"

	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}

func (us *userHandler) CreateUser(c *gin.Context) {
	var u dto.NewUserRequest

	if err := c.ShouldBindJSON(&u); err != nil {
		errBind := errs.NewUnprocessibleEntityError("invalid json request body")
		c.AbortWithStatusJSON(errBind.Status(), errBind)
		return
	}

	resp, err := us.userService.CreateUser(u)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (us *userHandler) Login(c *gin.Context) {
	var u dto.LoginRequest

	if err := c.ShouldBindJSON(&u); err != nil {
		errBind := errs.NewUnprocessibleEntityError("invalid json request body")
		c.AbortWithStatusJSON(errBind.Status(), errBind)
		return
	}

	resp, err := us.userService.Login(u)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (us *userHandler) Topup(c *gin.Context) {
	var u dto.TopUpRequest

	if err := c.ShouldBindJSON(&u); err != nil {
		errBind := errs.NewUnprocessibleEntityError("invalid json request body")
		c.AbortWithStatusJSON(errBind.Status(), errBind)
		return
	}

	jwtClaims := c.MustGet("user").(jwt.MapClaims)
	u.ID = int(jwtClaims["id"].(float64))

	resp, err := us.userService.Topup(u)
	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
