package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/adapter/input/handler"
	adapter "github.com/justahmed99/delos-farm/adapter/output"
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/usecase"

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

	db.AutoMigrate(&domain.Farm{}, &domain.Pond{})
	// db.AutoMigrate(&domain.Pond{})

	router := gin.Default()
	farmRepository := adapter.NewGormFarmRepository(db)
	farmHandler := handler.NewFarmHandler(usecase.NewFarmUsesCases(farmRepository))
	router.GET("/v1/farm/:id", farmHandler.GetFarmById)
	router.POST("/v1/farm", farmHandler.CreateFarm)
	router.PUT("/v1/farm/:id", farmHandler.UpdateFarm)
	router.DELETE("/v1/farm/:id", farmHandler.DeleteFarm)

	// pondHandler := handler.NewPondHandler(usecase.PondUseCases())
	// router.GET("/v1/farm/:id", pondHandler.GetFarmById)
	// router.POST("/v1/farm", pondHandler.CreateFarm)
	// router.PUT("/v1/farm/:id", pondHandler.UpdateFarm)
	// router.DELETE("/v1/farm/:id", pondHandler.DeleteFarm)

	router.Run("localhost:9090")
	// db.AutoMigrate(&domain.Farm{})
	// db.AutoMigrate(&domain.Pond{})
}
