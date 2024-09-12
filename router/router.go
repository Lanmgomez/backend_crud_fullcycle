package router

import (
	"net/http"
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
	router.POST("/login", Login)
	router.POST("/login/create-new-user", CreateNewUserLogin)
	router.GET("/login/:id", user.GetLoginLogsByUserID)

	return router
}

func CreateNewUserLogin(c *gin.Context) {

	var NewUser user.USERS

	if err := c.ShouldBindJSON(&NewUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao criar o usuário, dados inválidos",
		})
		return
	}

	if err := user.CreateNewUser(NewUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criar usuário no banco",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Usuário criado com sucesso",
	})
}

func Login(context *gin.Context) {

	var userLogin user.USERLOGIN

	clientIpAddress := user.FormattedIPAddress(context.ClientIP())

	if err := context.ShouldBindJSON(&userLogin); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos",
		})
		return
	}

	if err := user.SignIn(userLogin, clientIpAddress, context); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, true)
}
