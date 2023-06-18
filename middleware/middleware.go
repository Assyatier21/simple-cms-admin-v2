package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	cfg "github.com/assyatier21/simple-cms-admin-v2/config"
	m "github.com/assyatier21/simple-cms-admin-v2/models"
	"github.com/assyatier21/simple-cms-admin-v2/models/entity"
	"github.com/assyatier21/simple-cms-admin-v2/utils/constant"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var (
	authorization = "Authorization"
)

func getJWTSecretKey() string {
	return cfg.Load().JWTSecretKey
}

func ClaimToken(ctx echo.Context) entity.UserClaimsResponse {
	return ctx.Get("user").(entity.UserClaimsResponse)
}

func GenerateToken(registry entity.User) (t string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = registry.Name
	claims["phone"] = registry.Phone
	claims["role"] = registry.Role
	claims["expired_at"] = time.Now().Add(60 * time.Minute)

	t, err = token.SignedString([]byte(getJWTSecretKey()))
	if err != nil {
		log.Println("[Middleware] failed to signed jwt token, err: ", err)
		return
	}

	return
}

func ParseTokenJWT(tokenString string) (userClaims entity.UserClaimsResponse, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return []byte(getJWTSecretKey()), nil
	})
	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return userClaims, errors.New("invalid token claims")
	}

	expiredAtStr, _ := claims["expired_at"].(string)
	expiredAt, err := time.Parse(time.RFC3339, expiredAtStr)
	if err != nil {
		return userClaims, fmt.Errorf("failed to parse expired_at value: %v", err)
	}
	userClaims = entity.UserClaimsResponse{
		Name:      claims["name"].(string),
		Phone:     claims["phone"].(string),
		Role:      claims["role"].(string),
		ExpiredAt: expiredAt,
	}

	return
}

func Middleware(private bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var userClaims = entity.UserClaimsResponse{}
			tokenString := c.Request().Header.Get(authorization)

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("invalid token signing method")
				}
				return []byte(getJWTSecretKey()), nil
			})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, m.StandardResponse{Code: http.StatusUnauthorized, Status: constant.FAILED, Message: "invalid signing token method", Error: err})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, m.StandardResponse{Code: http.StatusUnauthorized, Status: constant.FAILED, Message: "invalid token claims", Error: errors.New("invalid token claims")})
			}

			userClaims.Name = claims["name"].(string)
			userClaims.Phone = claims["phone"].(string)
			userClaims.Role = claims["role"].(string)
			expiredAtStr, _ := claims["expired_at"].(string)
			expiredAt, _ := time.Parse(time.RFC3339, expiredAtStr)
			userClaims.ExpiredAt = expiredAt

			if IsPrivate(userClaims.Role, private) {
				return c.JSON(http.StatusUnauthorized, m.StandardResponse{Code: http.StatusUnauthorized, Status: constant.FAILED, Message: "invalid token claims", Error: errors.New("invalid token claims")})
			}

			if IsTokenExpired(expiredAt) {
				return c.JSON(http.StatusUnauthorized, m.StandardResponse{Code: http.StatusUnauthorized, Status: constant.FAILED, Message: "token already expired", Error: errors.New("token already expired")})
			}

			c.Set("user", userClaims)
			return next(c)
		}
	}
}

func IsPrivate(role string, private bool) bool {
	return (role != "admin" && private)
}

func IsTokenExpired(t time.Time) bool {
	now := time.Now()
	return t.Before(now)
}
