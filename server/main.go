package main

import (
	"fmt"
	"log"
	"os"
	"repos"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Props struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	setupRoutes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

func setupRoutes(app *fiber.App) {
	// app.Get("/get", returnScrapedData)
	app.Post("/api", returnPostedData)
}

func returnPostedData(c *fiber.Ctx) error {
	prop := &Props{}
	if err := c.BodyParser(&prop); err != nil {
		return err
	}
	s := []string{
		prop.Owner,
		prop.Name,
	}
	grade := repos.GetRepo(s)
	fmt.Println(grade)
	return c.JSON(grade)
}
