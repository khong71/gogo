package main

import (
	"gogo/database"
	"gogo/route"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func DB() {
	database.DatabaseInitMysql() // Connect to database SqlServer
	// database.DatabaseInitMysql() // Connect to  Database MySql
}

func main() {

	DB();

	// Initialize a new Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
	}))

	route.RouteInit(app)
	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
