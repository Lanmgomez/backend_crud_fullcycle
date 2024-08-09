package user

import (
	"database/sql"
	"log"
	"net/http"

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
	rows, err := db.Query("SELECT * FROM users")

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
