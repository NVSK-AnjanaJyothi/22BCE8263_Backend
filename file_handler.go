package handlers

import (
	"FileManager/config"
	"FileManager/models"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

// List all files
func ListFiles(c *fiber.Ctx) error {
	var files []models.File
	if err := config.DB.Find(&files).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch files"})
	}
	return c.JSON(files)
}

// Download a file
func DownloadFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	filePath := fmt.Sprintf("uploads/%s", filename)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return c.Status(404).JSON(fiber.Map{"error": "File not found"})
	}

	return c.SendFile(filePath)
}

// Delete a file
func DeleteFile(c *fiber.Ctx) error {
	id := c.Params("id")

	var file models.File
	if err := config.DB.First(&file, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "File not found"})
	}

	filePath := fmt.Sprintf("uploads/%s", file.Filename)
	os.Remove(filePath)

	config.DB.Delete(&file)

	return c.JSON(fiber.Map{"message": "✅ File deleted successfully"})
}
