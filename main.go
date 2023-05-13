package main

import (
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

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "https://opensourcegrader.vercel.app",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))

	app.Use(cors.New())

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

	return c.JSON(grade)
}
