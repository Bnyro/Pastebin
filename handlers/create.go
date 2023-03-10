package handlers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pastebin/entities"
)

var RootPath string

func Create(c *fiber.Ctx) error {
	entry := new(entities.Entry);
	if err := c.BodyParser(entry); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	entry.UUID = uuid.NewString()

	os.Mkdir(RootPath, 0755)
	err := os.WriteFile(RootPath + entry.UUID, []byte(entry.Content), 0755)
	if err != nil {
		panic(err)
	}

	return c.Render("created", entry)
}