package main

import( 
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the TutorLink API",
		})
	})
	
	r.Run(":8080")
}