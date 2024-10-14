package user

import (
	"errors"
	"fmt"
	"strconv"

	"crypto/rand"
	"encoding/hex"

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

func GenerateToken() (string, error) {
	// Cria um slice de 16 bytes (128 bits)
	bytes := make([]byte, 16)

	// Preenche o slice com bytes aleatórios
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Converte os bytes para uma string hexadecimal
	return hex.EncodeToString(bytes), nil
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

func GetAllPayments(c *gin.Context) ([]PAYMENT_METHOD, error) {
	id := c.Param("id")
	userPaymentID := parseParamIDtoInt(id)

	getPayments, err := GetAllPaymentsByUserIDInDB(userPaymentID)

	if err != nil {
		return nil, err
	}

	return getPayments, nil
}

func CreatePayment(c *gin.Context, payment PAYMENTS) error {

	token, errorToken := GenerateToken()
	if errorToken != nil {
		return errorToken
	}

	// Check if user exists before inserting
	existingUserID, err := CheckIfUserExists(c, payment.UserIdentification)
	if err != nil {
		return err
	}

	if existingUserID != 0 { // if exists, only insert new payment data
		if err := InsertNewPaymentInDB(payment.PaymentForm, existingUserID, token); err != nil {
			return err
		}

		return nil
	}

	if existingUserID == 0 { // if not exists, insert new user and then insert new payment data
		userID, err := InserUserIdentificationInDB(payment.UserIdentification) 
		if err != nil {
			return err
		}
	
		if err := InsertNewPaymentInDB(payment.PaymentForm, userID, token); err != nil {
			return err
		}
	
		return nil
	}

	return nil
}

func GetUfStatesList(c *gin.Context) ([]UF_STATES, error) {
	
	getUfStates, err := GetUfStatesInDB(c)

	if err != nil {
		return nil, err
	}

	return getUfStates, nil
}