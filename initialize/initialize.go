package initialize

import (
	"parameter-testing/config"
	log "parameter-testing/logger"
	"parameter-testing/utils"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Application struct {
	Config   *config.Config
	Database *gorm.DB
}

func InitApp() *Application {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := utils.GetEnv("ENV", "development")

	cfg := config.GetConfig(env)

	app := &Application{
		Config:   cfg,
		Database: initializeDB(cfg),
	}

	return app
}

func initializeDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.Database.Config()), &gorm.Config{})

	if err != nil {
		log.Fatalf("error initialize database : %s", err.Error())
	}

	log.Info("Database:\n",
		"config", cfg.Database.ConfigInfo(),
	)

	return db
}
