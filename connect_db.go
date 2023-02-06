package main

import (
	"github.com/justahmed99/delos-farm/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=development dbname=delos_farm port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	db.AutoMigrate(&domain.Farm{})
	db.AutoMigrate(&domain.Pond{})
}
