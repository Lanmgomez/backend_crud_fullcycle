package user

import (
	"database/sql"
	"log"
	"net/http"
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

func GetAllUsersInDB(c *gin.Context, activeUserStatus string) []USERSCRUD {
	rows, err := db.Query("SELECT * FROM crudusers WHERE activeUser = ?", activeUserStatus)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return users
		}

		users = append(users, user)
	}

	return users
}

// func GetUsers(c *gin.Context) {

// 	var activeUserStatus string = "ATIVO"

// 	rows, err := db.Query("SELECT * FROM crudusers WHERE activeUser = ?", activeUserStatus)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 	}

// 	defer rows.Close()

// 	var users []USERSCRUD

// 	for rows.Next() {
// 		var user USERSCRUD

// 		if err := rows.Scan(
// 			&user.ID,
// 			&user.Name,
// 			&user.Lastname,
// 			&user.Email,
// 			&user.Birthday,
// 			&user.Phone,
// 			&user.Address,
// 			&user.CreatedAt,
// 			&user.UpdatedAt,
// 			&user.ActiveUser,
// 		); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": err.Error(),
// 			})
// 			return
// 		}

// 		users = append(users, user)
// 	}

// 	c.JSON(http.StatusOK, users)
// }

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	parsedIDtoInt := parseParamIDtoInt(id)

	rows := db.QueryRow("SELECT * FROM crudusers WHERE id = ?", parsedIDtoInt)

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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	parsedIDtoInt := parseParamIDtoInt(id)

	var UpdateUser USERSCRUD

	if err := c.BindJSON(&UpdateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	UpdateUser.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	_, err := db.Exec(
		"UPDATE crudusers SET name = ?, lastname = ?, email = ?, birthday = ?, phone = ?, address = ?, updatedAt = ? WHERE id = ?",
		UpdateUser.Name,
		UpdateUser.Lastname,
		UpdateUser.Email,
		UpdateUser.Birthday,
		UpdateUser.Phone,
		UpdateUser.Address,
		UpdateUser.UpdatedAt,
		parsedIDtoInt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UpdateUser)
}

func CreateUser(c *gin.Context) {
	var CreateNewUser USERSCRUD
	var activeUserStatus string = "ATIVO"

	if err := c.BindJSON(&CreateNewUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	CreateNewUser.ActiveUser = activeUserStatus
	CreateNewUser.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	CreateNewUser.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	_, err := db.Exec(
		"INSERT INTO crudusers (name, lastname, email, birthday, phone, address, ActiveUser, createdAt, UpdatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		CreateNewUser.Name,
		CreateNewUser.Lastname,
		CreateNewUser.Email,
		CreateNewUser.Birthday,
		CreateNewUser.Phone,
		CreateNewUser.Address,
		CreateNewUser.ActiveUser,
		CreateNewUser.CreatedAt,
		CreateNewUser.UpdatedAt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, CreateNewUser)
}

func DeleteLogicalUserByID(c *gin.Context) {
	id := c.Param("id")
	parsedIDtoInt := parseParamIDtoInt(id)

	var logicDelete USERSCRUD
	var inactiveUser string = "INATIVO"

	if err := c.BindJSON(&logicDelete); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := db.Exec("UPDATE crudusers SET activeUser = ? WHERE id = ?",
		inactiveUser,
		parsedIDtoInt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, true)
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

func GetLoginLogsByUserID(c *gin.Context) {
	var loginLogs []LOGINLOGS

	id := c.Param("id")
	parsedIDtoInt := parseParamIDtoInt(id)

	rows, err := db.Query("SELECT * FROM loginLogs WHERE userId = ?",
		parsedIDtoInt,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		loginLogs = append(loginLogs, loginLog)
	}

	c.JSON(http.StatusOK, loginLogs)
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