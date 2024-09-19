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
	router.GET("/crudusers", GetCrudUsers)
	router.GET("/crudusers/:id", GetCrudUserByID)
	router.POST("/crudusers", CreateNewUserInCrud)
	router.PUT("/crudusers/:id", UpdateCrudUser)
	router.PATCH("/crudusers/:id", DeleteLogicalUserInCrud)

	// Login
	router.POST("/login", Login)
	router.POST("/login/create-new-user", CreateNewUserLogin)
	router.GET("/login/:id", GetLoginLogsByUserID)
	router.GET("/users", GetUsers)

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

func GetLoginLogsByUserID(c *gin.Context) {

	loginLogsRegisters, err := user.GetUserLoginDataByUserID(c)	

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, loginLogsRegisters)
}

func GetCrudUsers(c *gin.Context) {

	getUsers, err := user.GetUsers(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, getUsers)
}

func GetCrudUserByID(c *gin.Context) {

	getCrudByUserID, err := user.GetUserDataByID(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, getCrudByUserID)
}

func CreateNewUserInCrud(c *gin.Context) {

	var CreateNewUserData user.USERSCRUD

	if err := c.ShouldBindJSON(&CreateNewUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao criar o usuário, dados inválidos",
		})
		return
	}

	if err := user.NewUserInCrud(CreateNewUserData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func UpdateCrudUser(c *gin.Context) {

	var UpdateData user.USERSCRUD

	if err := c.ShouldBindJSON(&UpdateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao atualizar, dados inválidos",
		})
		return
	}

	if err := user.UpdatedUserInCrud(c, UpdateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func DeleteLogicalUserInCrud(c *gin.Context) {

	if err := user.DeleteLogicalUserByID(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

func GetUsers(c *gin.Context) {

	getUsers, err := user.GetLoginUsers(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, getUsers)
}