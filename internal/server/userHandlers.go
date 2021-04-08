package server

import "github.com/gin-gonic/gin"

func (s *Server) GetUserInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Get("user")
		if !ok {
			c.JSON(403, map[string]string{
				"error": "not authorized",
			})
			return
		}

		c.JSON(200, user)
		return
	}
}
