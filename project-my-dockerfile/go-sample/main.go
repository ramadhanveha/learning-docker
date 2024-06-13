package main

import (
	// "database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// get env
	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbName := os.Getenv("DB_NAME")
	// dbHost := os.Getenv("DB_HOST")
	dbPort := "78"

	// // connect database
	// connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	// fmt.Println("Connect to: ", connStr)

	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	fmt.Println("Connection to database failed: ", err)
	// }

	// // test ping to database
	// if err := db.Ping(); err != nil {
	// 	fmt.Printf("Ping to database failed (DB_HOST: %s): %s\n", dbHost, err)
	// } else {
	// 	fmt.Println("Successfully connected to DB : ", db)
	// }

	// create gin router
	r := gin.Default()

	// middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// define routes
	r.GET("/", func(c *gin.Context) {
		fmt.Println("Ouch")
		c.String(http.StatusOK, "Ouch")
	})

	httpPort := dbPort
	if httpPort == "" {
		httpPort = "78"
	}

	// Print port for debugging
	fmt.Println("Starting server on port:", httpPort)

	// start server
	if err := r.Run(":" + httpPort); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
