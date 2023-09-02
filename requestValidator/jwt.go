package validator

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/curiousz-peel/web-learning-platform-backend/config"
	userService "github.com/curiousz-peel/web-learning-platform-backend/service/user"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

var JWTSecret []byte

func InitSecretJWT() {
	jwtConfig, err := config.InitJWTConfig()
	if err != nil {
		log.Fatal("could not get JWT secret")
	}
	JWTSecret = []byte(jwtConfig.JWTSecret)
}

func GetLoginToken(userName string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	expirationTime := time.Now().Add(120 * time.Minute).Unix()
	claims["exp"] = expirationTime
	claims["authorized"] = true
	claims["iss"] = userName

	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", errors.New("could not sign token string")
	}

	return tokenString, nil
}

func ValidateToken(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) != 2 {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "failed to parse auth token",
			"data":    nil})
	}

	if bearerToken[0] != "Bearer" {
		return ctx.Status(http.StatusForbidden).JSON(&fiber.Map{
			"message": "auth token needs to be in the Bearer <token> format",
			"data":    nil})
	}
	token, err := checkToken(bearerToken[1])
	if err != nil {
		return ctx.Status(http.StatusForbidden).JSON(&fiber.Map{
			"message": "authorization failed",
			"data":    err.Error()})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		username, ok := claims["iss"].(string)
		if !ok {
			return ctx.Status(http.StatusForbidden).JSON(&fiber.Map{
				"message": "invalid auth token",
				"data":    nil})
		}

		_, err := userService.GetUserByUsername(username)
		if err != nil {
			return ctx.Status(http.StatusForbidden).JSON(&fiber.Map{
				"message": "failed to authenticate the user",
				"data":    nil})
		}
	}
	ctx.Set("iss", claims["iss"].(string))
	ctx.Next()
	return nil
}

func checkToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, errors.New("invalid token claims")
	}
	return token, nil
}
