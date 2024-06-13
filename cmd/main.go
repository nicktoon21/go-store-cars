package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	var psqlconn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	poolConfig, err := pgxpool.ParseConfig(psqlconn)
	CheckError(err)

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	CheckError(err)

	err = db.Ping(context.Background())
	CheckError(err)

	defer db.Close()
	fmt.Println("Connected!")

	eng := gin.Default()
	if err := eng.Run(); err != nil {
		panic(err)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
