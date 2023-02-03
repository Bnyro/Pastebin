package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/pastebin/database"
	"github.com/pastebin/handlers"
)

func main() {
	database.Init()

	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	});
	app.Use(cors.New())

	app.Static("/static", "./static")
	app.Get("/", handlers.Home)
	app.Post("/", handlers.Create)
	app.Get("/:id", handlers.Read)

	url := flag.String("addr", ":8080", "addr to listen on")
	flag.Parse()

	log.Fatal(app.Listen(*url))
}