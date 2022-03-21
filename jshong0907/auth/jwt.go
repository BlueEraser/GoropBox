package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorop-box/models"
	"gorop-box/services"
)

type JwtClaims struct {
	jwt.StandardClaims
	Email    string `json:"email"`
	NickName string `json:"NickName"`
}

func GetUserByJwt(c echo.Context) (*models.User, error) {
	userJwt := c.Get("user").(*jwt.Token)
	claims := userJwt.Claims.(*JwtClaims)
	return services.GetUserByEmail(claims.Email)
}
