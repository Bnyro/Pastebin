package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New();

	url := flag.String("addr", ":8080", "addr to listen on")
	flag.Parse()
	
	log.Fatal(app.Listen(*url))
}