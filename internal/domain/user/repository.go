package user

import (
	"database/sql"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

func InitDB(c *gin.Context) {
	dsn := "root:crudfullcycle@tcp(127.0.0.1:3307)/crudfullcycle"

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
}

func GetAllUsersInDB(c *gin.Context, activeUserStatus string) ([]USERSCRUD, error) {
	rows, err := db.Query("SELECT * FROM crudusers WHERE activeUser = ?", activeUserStatus)

	if err != nil {
		return nil, err
	}	

	defer rows.Close()

	var users []USERSCRUD

	for rows.Next() {
		var user USERSCRUD

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Lastname,
			&user.Email,
			&user.Birthday,
			&user.Phone,
			&user.Address,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.ActiveUser,
		); err != nil {
			return users, nil
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserDataCRUDByIDInDB(crudUserID int) (USERSCRUD, error) {
	rows := db.QueryRow("SELECT * FROM crudusers WHERE id = ?", crudUserID)

	var user USERSCRUD

	if err := rows.Scan(
		&user.ID,
		&user.Name,
		&user.Lastname,
		&user.Email,
		&user.Birthday,
		&user.Phone,
		&user.Address,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.ActiveUser,
	); err != nil {
		return user, err
	}

	return user, nil
}

func InsertNewUserCrudInDB(NewUserCrud USERSCRUD) error {

	_, err := db.Exec(
		"INSERT INTO crudusers (name, lastname, email, birthday, phone, address, ActiveUser, createdAt, UpdatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		NewUserCrud.Name,
		NewUserCrud.Lastname,
		NewUserCrud.Email,
		NewUserCrud.Birthday,
		NewUserCrud.Phone,
		NewUserCrud.Address,
		"ATIVO",
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02 15:04:05"),
	)

	if err != nil {
		return err
	}

	return nil
}

func UpdateUserCrudInDB(UpdateData USERSCRUD, crudUserIDToUpdate int) error {

	_, err := db.Exec(
		"UPDATE crudusers SET name = ?, lastname = ?, email = ?, birthday = ?, phone = ?, address = ?, updatedAt = ? WHERE id = ?",
		UpdateData.Name,
		UpdateData.Lastname,
		UpdateData.Email,
		UpdateData.Birthday,
		UpdateData.Phone,
		UpdateData.Address,
		time.Now().Format("2006-01-02 15:04:05"),
		crudUserIDToUpdate,
	)

	if err != nil {
		return err
	}

	return nil
}

func GetUserRegisteredInDB(userLogin string) (USERS, error) {
	var user USERS

	rows := db.QueryRow("SELECT id, userPassword FROM users WHERE username = ?", userLogin)

	if err := rows.Scan(&user.Id, &user.Password); err != nil {
		return USERS{}, err
	}

	return user, nil
}

func InsertLoginLogs(userId int, context *gin.Context, clientIpAddress string) error {

	_, err := db.Exec("INSERT INTO loginLogs (userId, userAgent, loginTime, status, ipAddress) VALUES (?, ?, ?, ?, ?)",
		userId,
		context.Request.UserAgent(),
		time.Now().Format("2006-01-02 15:04:05"), 
		"SUCCESS",
		clientIpAddress,
	)

	if err != nil {
		return err
	}

	return nil
}

func GetUserLoginDataByUserIDInDB(parsedIDtoInt int) ([]LOGINLOGS, error) {
	var loginLogs []LOGINLOGS

	rows, err := db.Query("SELECT * FROM loginLogs WHERE userId = ?",
		parsedIDtoInt,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var loginLog LOGINLOGS

		if err := rows.Scan(
			&loginLog.Id, 
			&loginLog.UserId, 
			&loginLog.UserAgent, 
			&loginLog.LoginTime, 
			&loginLog.Status, 
			&loginLog.IpAddress,
			); err != nil {
			return nil, err
		}
		loginLogs = append(loginLogs, loginLog)
	}

	return loginLogs, nil
}

func InsertNewUserInDB(CreateNewUser USERS) error {

	_, err := db.Exec(
		"INSERT INTO users (username, userPassword, createdAt, updatedAt) VALUES (?, ?, ?, ?)",
		CreateNewUser.Username,
		CreateNewUser.Password,
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02 15:04:05"),
	)

	if err != nil {
		return err
	}

	return nil
}

func DeleteLogicalUserInDB(IDtoLogicalDelete int) error {

	_, err := db.Exec("UPDATE crudusers SET activeUser = ? WHERE id = ?",
		"INATIVO",
		IDtoLogicalDelete,
	)

	if err != nil {
		return err
	}

	return nil
}

func GetLoginUsersInDB(c *gin.Context) ([]USERS, error) {

	var users []USERS

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user USERS

		if err := rows.Scan(
			&user.Id, 
			&user.Username, 
			&user.Password, 
			&user.CreatedAt, 
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetAllPaymentsByUserIDInDB(userPaymentID int) ([]PAYMENT_METHOD, error) {
	var payments []PAYMENT_METHOD

	rows, err := db.Query("SELECT * FROM userPaymentMethod WHERE paymentUserID = ?", userPaymentID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var payment PAYMENT_METHOD

		if err := rows.Scan(
			&payment.Id,
			&payment.PaymentUserId,
			&payment.PaymentFormInstallment,
			&payment.Token,
			&payment.DateTime,
		); err != nil {
			return nil, err
		}

		payments = append(payments, payment)
	}

	return payments, nil
}

func InserUserIdentificationInDB(userIdentification USER_IDENTIFICATION_CONTACT) (int64, error) {

	result, err := db.Exec("INSERT INTO userIdentificationContact (fullName, email, cpfOrCnpj, phone, address, addressNumber, neighborhood, city, uf, zipCode) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		userIdentification.FullName,
		userIdentification.Email,
		userIdentification.CpfOrCnpj,
		userIdentification.Phone,
		userIdentification.Address,
		userIdentification.AddressNumber,
		userIdentification.Neighborhood,
		userIdentification.City,
		userIdentification.Uf,
		userIdentification.ZipCode,
	)

	if err != nil {
		return 0, nil
	}

	// retornar id do user que foi inserido no BD
	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func InsertNewPaymentInDB(paymentForm PAYMENT_METHOD, userID int64, token string) error {

	_, err := db.Exec("INSERT INTO userPaymentMethod (paymentUserID, paymentFormInstallment, token, paymentDateTime) VALUES (?, ?, ?, ?)",
		userID,
		paymentForm.PaymentFormInstallment,
		token,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	if err != nil {
		return err
	}

	return nil
}

func CheckIfUserExists(context *gin.Context, userIdentification USER_IDENTIFICATION_CONTACT) (int64, error) {

	var userID int64

	query := "SELECT id FROM userIdentificationContact WHERE cpfOrCnpj = ? LIMIT 1"

	err := db.QueryRowContext(context, query, userIdentification.CpfOrCnpj).Scan(&userID)

	if err != nil {
		if err == sql.ErrNoRows {
			// User not found
			return 0, nil
		}
		// Other error occurred
		return 0, err
	}

	return userID, nil
}

func GetUfStatesInDB(c *gin.Context) ([]UF_STATES, error) {
	var ufStates []UF_STATES

	rows, err := db.Query("SELECT * FROM ufStatesList")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ufState UF_STATES

		if err := rows.Scan(
			&ufState.Id, 
			&ufState.State, 
			&ufState.Uf, 
		); err != nil {
			return nil, err
		}

		ufStates = append(ufStates, ufState)
	}

	return ufStates, nil	
}