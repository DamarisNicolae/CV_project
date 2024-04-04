package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	ID             int64
	Jobtitle       string
	Firstname      string
	Lastname       string
	Email          string
	Phone          string
	Adress         string
	City           string
	Country        string
	Postalcode     string
	Dateofbirth    string
	Nationality    string
	Summary        string
	Workexperience string
	Education      string
	Skills         string
	Languages      string
}

func userByID(id int64) (User, error) {
	var user User

	row := db.QueryRow("SELECT * FROM user WHERE id = ?", id)
	if err := row.Scan(&user.ID, &user.Jobtitle, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Adress, &user.City, &user.Country, &user.Postalcode, &user.Dateofbirth, &user.Nationality, &user.Summary, &user.Workexperience, &user.Education, &user.Skills, &user.Languages); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("userById %d: no such user", id)
		}
		return user, fmt.Errorf("userById %d: %v", id, err)
	}
	return user, nil
}

func main() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "users",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN()) // to initialize the db variable, passing the return value of FormatDSN.
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping() // to confirm that the database/sql package can connect
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	user, err := userByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User found: %v\n", user)

}
