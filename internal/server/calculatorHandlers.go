package server

import (
	"encoding/json"

	"github.com/estromenko/yoof-api/internal/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetCalories() gin.HandlerFunc {
	type request struct {
		EatingAmount int     `json:"eating_amount"`
		Gender       int     `json:"gender"`
		Weight       float32 `json:"weight"`
		Growth       float32 `json:"growth"`
		Training     int     `json:"training"`
	}
	return func(c *gin.Context) {
		var req request
		json.NewDecoder(c.Request.Body).Decode(&req)

		calories := s.Services().CalcService.CountCalories(req.Gender, req.Weight, req.Growth, req.Training)
		dayList := s.Services().CalcService.GetDayList(req.EatingAmount, calories)
		if len(dayList) == 0 {
			c.JSON(400, map[string]string{
				"error": "wrong eating amount",
			})
			return
		}

		var dishes []*models.Dish

		for i, v := range dayList {
			if v != 0.0 {
				dish := s.Services().CalcService.GetRandomDishByCalories(dayList[i], i, 40)
				dishes = append(dishes, dish)
			}
		}

		c.JSON(200, dishes)
	}
}
