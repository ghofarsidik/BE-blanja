package main

import (
	"log"
	"os"
	"server/src/configs"
	"server/src/helpers"
	"server/src/routes"
	"server/src/services"

	"github.com/gofiber/helmet/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return "0.0.0.0:" + port
}

func main() {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	app := fiber.New()

	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowMethods:  "GET,POST,PUT,DELETE",
		AllowHeaders:  "*",
		ExposeHeaders: "Content-Length",
	}))

	configs.InitDB()
	services.InitMidtrans()
	helpers.Migration()
	routes.Router(app)

	// Log error if app fails to listen
	if err := app.Listen(getPort()); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
