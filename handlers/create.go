package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pastebin/database"
	"github.com/pastebin/entities"
)

func Create(c *fiber.Ctx) error {
	entry := new(entities.Entry);
	if err := c.BodyParser(entry); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	entry.UUID = uuid.NewString()
	database.Db.Create(entry)
	return c.Render("created", entry)
}