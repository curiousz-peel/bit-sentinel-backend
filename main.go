package main

import (
	mail "github.com/curiousz-peel/web-learning-platform-backend/mailer"
	"github.com/curiousz-peel/web-learning-platform-backend/routes"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/gofiber/fiber/v2"
)

func main() {

	storage.ConnectDb()
	mail.InitMail()

	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":8080")
}
