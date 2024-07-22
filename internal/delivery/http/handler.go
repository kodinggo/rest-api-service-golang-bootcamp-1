package handler

import (
	"kodinggo/internal/config"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type jwtCustomClaims struct {
	Id       int    `json:"id"`
	Username string `json:"name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func jwtConfig() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(config.GetJwtSecret()),
	}
}

func signJwtToken(id int, username string, role string) (string, error) {
	claims := &jwtCustomClaims{
		id,
		username,
		role,
		jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.GetJwtSecret()))
	if err != nil {
		return "", err
	}

	return t, nil
}

func claimSession(c echo.Context) *jwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)

	return claims
}
