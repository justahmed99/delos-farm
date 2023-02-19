package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/adapter/input/handler"
	adapter "github.com/justahmed99/delos-farm/adapter/output"
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/usecase"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to delos-farm API")
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.Farm{}, &domain.Pond{}, &domain.Agent{}, &domain.RequestRecord{})

	router := gin.Default()
	farmRepository := adapter.NewGormFarmRepository(db)
	farmHandler := handler.NewFarmHandler(usecase.NewFarmUsesCases(farmRepository))
	router.GET("/v1/farm/:id", farmHandler.GetFarmById)
	router.GET("/v1/farm", farmHandler.GetFarms)
	router.POST("/v1/farm", farmHandler.CreateFarm)
	router.PUT("/v1/farm", farmHandler.UpdateFarm)
	router.DELETE("/v1/farm/:id", farmHandler.DeleteFarm)

	pondRepository := adapter.NewGormPondRepository(db)
	pondHandler := handler.NewPondHandler(usecase.NewPondUseCases(pondRepository))
	router.GET("/v1/pond/:id", pondHandler.GetPondById)
	router.GET("/v1/pond", pondHandler.GetPonds)
	router.GET("/v1/pond/farm/:farm_id", pondHandler.GetPondsByFarmID)
	router.POST("/v1/pond", pondHandler.CreatePond)
	router.PUT("/v1/pond", pondHandler.UpdatePond)
	router.DELETE("/v1/pond/:id", pondHandler.DeletePond)
	router.DELETE("/v1/pond/farm/:farm_id", pondHandler.DeletePondsByFarmID)

	monitorRepository := adapter.NewGormMonitorRepository(db)
	monitorHandler := handler.NewMonitorHandler(usecase.NewMonitorUseCases(monitorRepository))
	router.GET("/v1/monitor", monitorHandler.GetMonitorData)

	router.Run(appHost + ":" + appPort)
}
