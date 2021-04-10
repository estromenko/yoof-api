package server

import (
	"encoding/json"
	"strconv"

	"github.com/estromenko/yoof-api/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreateDish() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dish models.Dish
		if err := json.NewDecoder(c.Request.Body).Decode(&dish); err != nil {
			c.JSON(400, map[string]string{
				"error": "error decoding dish: " + err.Error(),
			})
			return
		}

		if err := s.Services().DishService.Create(&dish); err != nil {
			c.JSON(400, map[string]string{
				"error": "error decoding dish: " + err.Error(),
			})
			return
		}

		c.JSON(201, dish)
	}
}

func (s *Server) GetAllDishes() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "?"))
		offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

		if limit == '?' || limit == 0 {
			c.JSON(400, map[string]string{
				"error": "limit must be provided as query param",
			})
			return
		}

		dishes, err := s.Services().DishService.Repo().GetAll(limit, offset)
		if err != nil {
			c.JSON(400, map[string]string{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, dishes)
	}
}

func (s *Server) GetDish() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, map[string]string{
				"error": err.Error(),
			})
			return
		}

		dish, err := s.Services().DishService.Repo().FindByID(id)
		if err != nil {
			c.JSON(404, map[string]string{
				"error": "dish not found",
			})
			return
		}

		c.JSON(200, dish)
	}
}

func (s *Server) DeleteDish() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, map[string]string{
				"error": err.Error(),
			})
			return
		}

		if err := s.Services().DishService.Repo().DeleteByID(id); err != nil {
			c.JSON(400, map[string]string{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, map[string]string{
			"success": "dish deleted successfully",
		})
	}
}

func (s *Server) UpdateDish() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
