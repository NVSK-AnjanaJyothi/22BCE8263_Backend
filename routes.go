package routes

import (
	"FileManager/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the routes for the server
func RegisterRoutes(router *gin.Engine) {
	router.POST("/upload", handlers.UploadFileHandler) // ✅ For uploading files
	router.GET("/files", handlers.GetFileHandler)      // ✅ For listing files
	router.GET("/download/:filename", handlers.DownloadFileHandler)
	router.POST("/signup", handlers.Signup)
}
