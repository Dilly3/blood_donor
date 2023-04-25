package web

import "github.com/gin-gonic/gin"

func (c *Config) Test() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, "Welcome back")
	}
}
