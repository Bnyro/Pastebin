package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pastebin/database"
	"github.com/pastebin/entities"
)

func Read(c *fiber.Ctx) error {
	var entry entities.Entry
	uuid := c.Params("uuid")
	err := database.Db.First(&entry, "uuid = ?", uuid).Error
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.Render("read", entry)
}