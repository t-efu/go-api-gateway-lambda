package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	mysqlRepo "github.com/t-efu/go-api-gateway-lambda/infrastructure/persistence/mysql"
	"github.com/t-efu/go-api-gateway-lambda/interfaces/handler"
	"github.com/t-efu/go-api-gateway-lambda/usecase"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	endpoint = os.Getenv("DSN")
)

func main() {
	gormDB, err := gorm.Open(mysql.Open(endpoint), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed DB connection:", err)
	}
	db, err := gormDB.DB()
	if err != nil {
		log.Fatalln("failed create ")
	}
	defer db.Close()

	memoRepo := mysqlRepo.NewMemoRepository(gormDB)
	memoUsecase := usecase.NewMemoUsecase(memoRepo)
	memoHandler := handler.NewMemoHandler(memoUsecase)
	lambda.Start(memoHandler.Find)
}
