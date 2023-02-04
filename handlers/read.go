package handlers

import (
	"html/template"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/pastebin/utilities"
)

func Read(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	data, err := os.ReadFile(RootPath + uuid)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	content := utilities.Format(string(data))
	return c.Render("read", fiber.Map {
		"Content": template.HTML(content),
	})
}