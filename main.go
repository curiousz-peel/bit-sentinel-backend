package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/curiousz-peel/web-learning-platform-backend/models"
	"github.com/curiousz-peel/web-learning-platform-backend/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Author struct {
	ID          int      `json:"id"`
	User        User     `json:"user"`
	Profession  string   `json:"profession"`
	Description string   `json:"description"`
	Topics      []string `json:"topics"`
}

type Media struct {
	ID        int       `json:"id"`
	CourseID  int       `json:"courseID"`
	FilePath  string    `json:"filePath"`
	FileType  string    `json:"fileType"`
	CreatedAt time.Time `json:"createdAt"`
}

type Course struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author Author `json:"author"`
	//need to decide on the Content type..
	Content   []Media   `json:"content"`
	Tags      []string  `json:"tags"`
	AddedDate time.Time `json:"addedDate"`
	Visible   bool      `json:"visible"`
	Rating    float32   `json:"rating"`
}

type Subscription struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	ExpiresAt time.Time `json:"expiresAt"`
	Price     int       `json:"price"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Birthday  time.Time `json:"birthday"`
	IsMod     bool      `json:"isModerator"`
}

type Account struct {
	ID               int          `json:"id"`
	User             User         `json:"user"`
	Subscription     Subscription `json:"subscription"`
	Courses          []Course     `json:"courses"`
	RegistrationDate time.Time    `json:"registrationDate"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateUser(ctx *fiber.Ctx) error {
	user := User{}
	err := ctx.BodyParser(&user)
	if err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&user).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not create user"})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "create user succeeded"})
	return nil
}

func (r *Repository) DeleteUserByID(ctx *fiber.Ctx) error {
	userModel := models.Users{}
	id := ctx.Params("id")
	if id == "" {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "user ID cannot be empty on delete"})
	}
	err := r.DB.Delete(userModel, id).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete user"})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{"message": "user deleted successfully"})
	return nil
}

func (r *Repository) GetUserByID(ctx *fiber.Ctx) error {
	userModel := &models.Users{}
	id := ctx.Params("id")
	if id == "" {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "user ID cannot be empty on get"})
	}
	fmt.Println("the user id is", id)
	err := r.DB.Where("id = ?", id).First(userModel).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not find the user"})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user found successfully",
		"data":    userModel})
	return nil
}

func (r *Repository) GetUsers(ctx *fiber.Ctx) error {
	userModels := &[]models.Users{}

	err := r.DB.Find(userModels).Error
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not fetch users"})
		return err
	}

	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "users fetched successfully",
		"data":    userModels})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/createUser", r.CreateUser)
	api.Delete("/deleteUser/:id", r.DeleteUserByID)
	api.Get("/getUser/:id", r.GetUserByID)
	api.Get("/users", r.GetUsers)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database")
	}

	err = models.MigrateUsers(db)
	if err != nil {
		log.Fatal("could not migrate the dB")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
