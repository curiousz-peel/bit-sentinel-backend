package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/curiousz-peel/web-learning-platform-backend/config"
	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	config, err := config.InitDbConfig()
	if err != nil {
		log.Fatal("could not initialize config")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connect to the database! \n", err.Error())
		os.Exit(2)
	}
	log.Println("connected to the database successfully")
	log.Println("starting migrations")
	db.AutoMigrate(&models.User{},
		&models.Author{},
		&models.Subscription{},
		&models.SubscriptionPlan{},
		&models.MediaType{},
		&models.Media{},
		&models.Comment{},
		&models.Course{},
		&models.Lesson{},
		&models.Enrollment{},
		&models.Option{},
		&models.Progress{},
		&models.Question{},
		&models.Quiz{},
		&models.Rating{})

	//pass the created db connection to the global DB variable
	DB = db
}
