package server

import (
	"encoding/json"

	"github.com/estromenko/yoof-api/internal/models"
	"github.com/gin-gonic/gin"
)

type (
	authResponseEntity struct {
		Token string `json:"token"`
		ID    int    `json:"id"`
	}
)

// @Summary List auth
// @Description Registration page
// @Accept  json
// @Produce  json
// @Router /auth/reg [post]
// @Success 201 {object} authResponseEntity
func (s *Server) RegistrationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		json.NewDecoder(c.Request.Body).Decode(&user)

		if err := s.services.UserService.Create(&user); err != nil {
			c.JSON(400, map[string]string{
				"error": "user creation: " + err.Error(),
			})
			return
		}

		token, err := s.services.UserService.GenerateToken(&user)
		if err != nil {
			if err := s.services.UserService.Create(&user); err != nil {
				c.JSON(400, map[string]string{
					"error": "token generation: " + err.Error(),
				})
				return
			}
		}

		c.JSON(201, map[string]interface{}{
			"token": token,
			"id":    user.ID,
		})
		return
	}
}

// @Summary List auth
// @Description Login page
// @Accept  json
// @Produce  json
// @Router /auth/login [post]
// @Success 201 {object} authResponseEntity
func (s *Server) LoginHandler() gin.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(c *gin.Context) {
		var req request
		json.NewDecoder(c.Request.Body).Decode(&req)

		user, token, err := s.services.UserService.Login(req.Email, req.Password)
		if err != nil {
			c.JSON(403, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		c.JSON(201, map[string]interface{}{
			"token": token,
			"id":    user.ID,
		})
		return
	}
}

