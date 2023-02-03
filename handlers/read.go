package handlers

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/pastebin/database"
	"github.com/pastebin/entities"
	"github.com/pastebin/utilities"
)

func Read(c *fiber.Ctx) error {
	var entry entities.Entry
	uuid := c.Params("uuid")
	err := database.Db.First(&entry, "uuid = ?", uuid).Error
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	content := utilities.Format(entry.Content)
	return c.Render("read", fiber.Map {
		"Content": template.HTML(content),
	})
}