package user

import (
	"errors"
	"fmt"
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

func GetUsers(c *gin.Context) ([]USERSCRUD, error) {
	var activeUserStatus string = "ATIVO"

	getAllUsers, err := GetAllUsersInDB(c, activeUserStatus)

	if err != nil {
		return nil, err
	}

	return getAllUsers, nil
}

func GetUserLoginDataByUserID(c *gin.Context) ([]LOGINLOGS, error) {
	id := c.Param("id")
	parsedIDtoInt := parseParamIDtoInt(id)

	result, err := GetUserLoginDataByUserIDInDB(parsedIDtoInt)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetUserDataByID(c *gin.Context) (USERSCRUD, error) {
	id := c.Param("id")
	crudUserID := parseParamIDtoInt(id)

	result, err := GetUserDataCRUDByIDInDB(crudUserID)

	if err != nil {
		return USERSCRUD{}, err
	}

	return result, nil
}

func NewUserInCrud(NewUserCrud USERSCRUD) error {
	if err := InsertNewUserCrudInDB(NewUserCrud); err != nil {
		return err
	}

	return nil
}

func UpdatedUserInCrud(c *gin.Context, UpdateData USERSCRUD) error {
	id := c.Param("id")
	crudUserIDToUpdate := parseParamIDtoInt(id)

	if err := UpdateUserCrudInDB(UpdateData, crudUserIDToUpdate); err != nil {
		return err
	}

	return nil
}

func DeleteLogicalUserByID(c *gin.Context) error {
	id := c.Param("id")
	IDtoLogicalDelete := parseParamIDtoInt(id)

	if err := DeleteLogicalUserInDB(IDtoLogicalDelete); err != nil {
		return err
	}

	return nil
}

func GetLoginUsers(c *gin.Context) ([]USERS, error) {

	getUsersLogin, err :=  GetLoginUsersInDB(c)

	if err != nil {
		return nil, err
	}

	return getUsersLogin, nil
}