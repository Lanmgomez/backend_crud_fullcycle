package router

import (
	"time"

	"github.com/Lanmgomez/backend_crud_fullcycle/internal/domain/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	// Users crud
	router.GET("/users", user.GetUsers)
	router.GET("/users/:id", user.GetUserByID)
	router.POST("/users", user.CreateUser)
	router.PUT("/users/:id", user.UpdateUser)
	router.PATCH("/users/:id", user.DeleteLogicalUserByID)

	// Login
	router.POST("/login", user.LoginHandler)
	router.GET("/login/:id", user.GetLoginLogsByUserID)

	return router
}
