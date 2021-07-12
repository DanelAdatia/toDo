package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/praveeenrm/todoapp/controllers"
)

// SetUpRoutes returns the routes
func SetUpRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	api := r.Group("/api")
	{
		api.GET("todos", controllers.GetAllTodosHandler)
		api.GET("todos/:id", controllers.GetATodoHandler)
		api.POST("todos", controllers.CreateATodoHandler)
		api.PUT("todos/:id", controllers.UpdateATodoHandler)
		api.DELETE("todos/:id", controllers.DeleteATodoHandler)
	}

	return r

}
