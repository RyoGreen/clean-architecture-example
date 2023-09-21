package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
	return
	db, err := sql.Open("postgres", "user=postgres dbname=postgres password=postgres host=postgres port=5432 sslmode=disable")
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Println(err)
		return
	}
	fmt.Println("done")
}
