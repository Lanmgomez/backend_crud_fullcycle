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

func InsertPaymentInDB(c *gin.Context, payment PAYMENTS, parsedIDtoInt int) error {

	_, err := db.Exec("INSERT INTO payments (paymentID, userPay, status, token, paymentDate) VALUES (?, ?, ?, ?, ?)",
		parsedIDtoInt,
		payment.UserPay,
		payment.Status,
		payment.Token,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	if err != nil {
		return err
	}

	return nil
}

func GetAllPaymentsByUserIDInDB(userPaymentID int) ([]PAYMENTS, error) {
	var payments []PAYMENTS

	rows, err := db.Query("SELECT * FROM payments WHERE userPaymentID = ?", userPaymentID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var payment PAYMENTS

		if err := rows.Scan(
			&payment.Id,
			&payment.UserPaymentID,
			&payment.UserPay,
			&payment.Status,
			&payment.Token,
			&payment.PaymentDate,
		); err != nil {
			return nil, err
		}

		payments = append(payments, payment)
	}

	return payments, nil
}