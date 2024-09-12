package user

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
func parseParamIDtoInt(id string) int {
	parsedID, err := strconv.ParseInt(id, 10, 64) // 10 base, 64 bits

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(parsedID)
}

func FormattedIPAddress(IpAddress string) string {
	if IpAddress == "::1" {
		return "127.0.0.1"
	}

	return IpAddress
}
func CreateNewUser(NewUser USERS) error {
	if err := InsertNewUserInDB(NewUser); err != nil {
		return err
	}

	return nil
}

func SignIn(userLogin USERLOGIN, clientIpAddress string, context *gin.Context) error {
	
	user, err := GetUserRegisteredInDB(userLogin.Username)

	if err != nil {
		return errors.New("usuário inválido")
	}		

	if userLogin.Password != user.Password {
		return errors.New("senha inválida")
	}

	if err := InsertLoginLogs(user.Id, context, clientIpAddress); err != nil {
		return errors.New("erro ao registrar o login")
	}

	return nil
}

func GetUsers(c *gin.Context) {
	var activeUserStatus string = "ATIVO"

	getAllUsers := GetAllUsersInDB(c, activeUserStatus)

	c.JSON(http.StatusOK, getAllUsers)
}