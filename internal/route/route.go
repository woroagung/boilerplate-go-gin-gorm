package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadRoute(r *gin.Engine) {
	// health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// other route here
	exampleRoute := r.Group("/example")
	exampleRoute.GET("/my_handler", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
}
