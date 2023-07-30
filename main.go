package main

import (
	mail "github.com/curiousz-peel/web-learning-platform-backend/mailer"
	"github.com/curiousz-peel/web-learning-platform-backend/routes"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/gofiber/fiber/v2"
)

// ORDER OF IMPLEMENTATION: Author, Media, Course, Account

// type Media struct {
// 	ID        int       `json:"id"`
// 	CourseID  int       `json:"courseID"`
// 	FilePath  string    `json:"filePath"`
// 	FileType  string    `json:"fileType"`
// 	CreatedAt time.Time `json:"createdAt"`
// }

// type Course struct {
// 	ID     int    `json:"id"`
// 	Title  string `json:"title"`
// 	Author []Author `json:"author"`
// 	//need to decide on the Content type..
// 	Content   []Media   `json:"content"`
// 	Tags      []string  `json:"tags"`
// 	AddedDate time.Time `json:"addedDate"`
// 	Visible   bool      `json:"visible"`
// 	Rating    float32   `json:"rating"`
// }

// type Account struct {
// 	ID               int          `json:"id"`
// 	User             User         `json:"user"`
// 	Subscription     Subscription `json:"subscription"`
//  SubscriptionExpiry time.Time  `json:"subscriptionExpiry`
// 	Courses          []Course     `json:"courses"`
// 	RegistrationDate time.Time    `json:"registrationDate"`
// }

func main() {

	storage.ConnectDb()
	mail.InitMail()

	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":8080")
}
