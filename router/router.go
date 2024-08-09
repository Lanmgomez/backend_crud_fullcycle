package router

import (
	"github.com/Lanmgomez/backend_crud_fullcycle/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()

	router.GET("/users", user.GetUsers)

	return router
}
