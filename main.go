package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:0341@tcp(localhost:3306)/mylab")
	if err != nil {
		panic(err)
	}
	err = MigrateDB(db)
	if err != nil {
		panic(err)
	}

	userRepo := NewUserRepo(db)

	usePointAsDiscountHandler := NewUsePointsAsDiscountHandler(userRepo)

	handler := NewHttpHandler(usePointAsDiscountHandler)

	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
