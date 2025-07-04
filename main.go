package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	  "github.com/shravanimaity/golangtut/queries"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

const (
	HOST     = "dpg-d1ibu4ali9vc73fpqjt0-a.oregon-postgres.render.com"
	PORT     = "5432"
	USERNAME = "bootcamp_db_yj7k_user"
	PASSWORD = "7uKulyPCAXk4pvvoLndAJb8mo0vExXYE"
	DBNAME   = "bootcamp_db_yj7k"
)

func GetPsqlInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		HOST, PORT, USERNAME, PASSWORD, DBNAME)
}

func CreateDbObject() error {
	var err error

	DB, err = sql.Open("postgres", GetPsqlInfo())
	if err != nil {
		return fmt.Errorf("error opening DB: %w", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %w", err)
	}

	fmt.Println("Connected successfully")

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxIdleTime(10 * time.Minute)
	DB.SetConnMaxLifetime(1 * time.Hour)

	return nil
}

func main() {
	if err := CreateDbObject(); err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer DB.Close()

	tx, err := DB.Begin()
	if err != nil {
		log.Fatal("TX error:", err)
	}
	defer tx.Commit()

	var pointerErr error
	users := queries.FetchAllUsersQuery(tx, &pointerErr)
	if pointerErr != nil {
		log.Println("Error fetching users:", pointerErr)
		return
	}

	for _, u := range users {
		fmt.Printf("User: ID=%d, Email=%s, Name=%s\n", u.ID, u.Email, getSafeName(u.Name))
	}
}

func getSafeName(name *string) string {
	if name == nil {
		return "(null)"
	}
	return *name
}
