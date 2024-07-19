package handler

import (
	"errors"
	"kodinggo/internal/config"
	"kodinggo/internal/model"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase model.IUserUsecase
}

func NewUserHandler(e *echo.Group, userUsecase model.IUserUsecase) {
	userHandler := &UserHandler{
		userUsecase: userUsecase,
	}

	e.POST("/users", userHandler.Create)
	e.POST("/users/login", userHandler.Login)
}

func (u *UserHandler) Create(c echo.Context) error {
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Message: err.Error(),
		})
	}

	err = u.userUsecase.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response{
		Status:  http.StatusCreated,
		Message: "success",
	})
}

func (u *UserHandler) Login(c echo.Context) error {
	var reqUser model.User
	err := c.Bind(&reqUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response{
			Message: err.Error(),
		})
	}

	user, err := u.userUsecase.Login(reqUser.Username, reqUser.Password)
	if err != nil {
		if errors.Is(err, model.ErrInvalidPassword) {
			return c.JSON(http.StatusUnauthorized, response{
				Status:  http.StatusUnauthorized,
				Message: err.Error(),
			})
		}

		if errors.Is(err, model.ErrUsernameNotFound) {
			return c.JSON(http.StatusNotFound, response{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			})
		}

		return c.JSON(http.StatusInternalServerError, response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	claims := &jwtCustomClaims{
		user.Id,
		user.Username,
		jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.GetJwtSecret()))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "success",
		Data:    map[string]string{"token": t},
	})
}
