package user

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

func GetUsers(c *gin.Context) {

	var activeUserStatus string = "ATIVO"

	rows, err := db.Query("SELECT * FROM users WHERE activeUser = ?", activeUserStatus)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	defer rows.Close()

	var users []USERS

	for rows.Next() {
		var user USERS

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

		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func parseParamIDtoInt(id string) int {
	parsedID, err := strconv.ParseInt(id, 10, 64) // 10 base, 64 bits

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(parsedID)
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	parsedIDtoInt := parseParamIDtoInt(id)

	rows := db.QueryRow("SELECT * FROM users WHERE id = ?", parsedIDtoInt)

	var user USERS

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

	var UpdateUser USERS

	if err := c.BindJSON(&UpdateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	UpdateUser.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	_, err := db.Exec(
		"UPDATE users SET name = ?, lastname = ?, email = ?, birthday = ?, phone = ?, address = ?, updatedAt = ? WHERE id = ?",
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
	var CreateNewUser USERS

	if err := c.BindJSON(&CreateNewUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	CreateNewUser.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	CreateNewUser.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	_, err := db.Exec(
		"INSERT INTO users (name, lastname, email, birthday, phone, address, createdAt, UpdatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		CreateNewUser.Name,
		CreateNewUser.Lastname,
		CreateNewUser.Email,
		CreateNewUser.Birthday,
		CreateNewUser.Phone,
		CreateNewUser.Address,
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

	var logicDelete USERS
	var inactiveUser string = "INATIVO"

	if err := c.BindJSON(&logicDelete); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := db.Exec("UPDATE users SET activeUser = ? WHERE id = ?",
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
