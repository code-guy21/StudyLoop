package app 

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/spf13/viper"
	"github.com/code-guy21/TutorLink/server/internal/handlers"
	"github.com/code-guy21/TutorLink/server/internal/domain/models"
)

type App struct {
	router *gin.Engine
	db *gorm.DB
}

func (a *App) setupRoutes(){
	a.router.GET("/health", handlers.HealthCheck)
}

func NewApp() (*App, error){
	if err := initConfig(); err != nil{
		return nil, fmt.Errorf("error initializing config: %w", err)
	}

	db, err := initDB()

	if err != nil {
		return nil, fmt.Errorf("error initializing database: %w", err)
	}

	router := gin.Default()

	app := &App {
		router: router,
		db: db,
	}

	app.setupRoutes()

	return app, nil
}

func (a *App) Run() error {
	return a.router.Run(fmt.Sprintf(":%d", viper.GetInt("server.port")))
}



func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	return nil
}

func initDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("database.host"),
		viper.GetInt("database.port"),
		viper.GetString("database.user"),	
		viper.GetString("database.password"),
		viper.GetString("database.name"), viper.GetString("database.sslmode"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	if err := db.AutoMigrate(&models.Session{}, &models.Payment{}); err != nil{
		return nil, fmt.Errorf("error migrating database: %w", err)
	}

	return db, nil
}