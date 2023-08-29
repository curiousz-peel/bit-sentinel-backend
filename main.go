package main

import (
	mail "github.com/curiousz-peel/web-learning-platform-backend/mailer"
	"github.com/curiousz-peel/web-learning-platform-backend/routes"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	storage.ConnectDb()
	mail.InitMail()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
	}))

	routes.SetupRoutes(app)
	app.Listen(":8080")
}
