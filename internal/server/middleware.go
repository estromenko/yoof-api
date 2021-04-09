package server

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func (s *Server) baseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.logger.Info().Str("url", c.Request.URL.Path).Str("method", c.Request.Method).Send()
	}
}

func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var unprocessedToken string

		head := string(c.Request.Header.Get("Authorization"))
		splited := strings.Split(head, " ")
		if len(splited) != 2 {
			c.Abort()
			c.JSON(403, map[string]interface{}{
				"error": "Token is not provided",
			})
			return
		}
		unprocessedToken = splited[1]

		token, err := jwt.Parse(unprocessedToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method, got " + t.Header["alg"].(string))
			}

			return []byte(s.services.UserService.Config.JWTSecret), nil
		})

		if err != nil {
			c.Abort()
			c.JSON(403, map[string]string{
				"error": "token error: " + err.Error(),
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		id := claims["id"].(float64)
		user, err := s.Services().UserService.Repo().FindByID(int(id))
		if err != nil {
			c.Abort()
			c.JSON(403, map[string]string{
				"error": "token error: " + err.Error(),
			})
			return
		}

		c.Set("user", user)
	}
}
