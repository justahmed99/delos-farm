package main

import (
	"fmt"

	"github.com/justahmed99/delos-farm/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Print("Welcome to delos-farm API")

	dsn := "host=localhost user=postgres password=development dbname=delos_farm port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.Farm{})
}
