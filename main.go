package main

import (
	"github.com/Lanmgomez/backend_crud_fullcycle/internal/domain/user"

	"github.com/Lanmgomez/backend_crud_fullcycle/router"
	"github.com/gin-gonic/gin"
)

func main() {
	c := &gin.Context{}
	user.InitDB(c)

	r := router.Routers()
	r.Run(":5000")
}
