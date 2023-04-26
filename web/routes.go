package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Config) Serve() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/save", c.HandleSaveCandidate())
	router.GET("/all", c.HandleGetAllCandidates())
	router.GET("/name/:fullname", c.HandleGEtByNAme())
	router.GET("/:id", c.HandleGEtById())

	return router
}
