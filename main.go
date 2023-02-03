package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pastebin/database"
	"github.com/pastebin/handlers"
)

func main() {
	database.Init()

	app := fiber.New();
	app.Use(cors.New())

	app.Get("/", handlers.Home)
	app.Post("/", handlers.Create)
	app.Get("/:id", handlers.Read)

	url := flag.String("addr", ":8080", "addr to listen on")
	flag.Parse()

	log.Fatal(app.Listen(*url))
}