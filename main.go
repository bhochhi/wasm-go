package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//host static
	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	//host api
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ping": "pong",
			})
		})
	}

	router.Run(":3000")

}
