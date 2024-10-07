package main

import( 
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/code-guy21/TutorLink/server/internal/repositories"
)

func main(){
	repositories.InitDatabse()

	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the TutorLink API",
		})
	})
	
	r.Run(":8080")
}