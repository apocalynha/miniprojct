package service

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func CloudinaryUpload(c echo.Context, fileheader string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("gagal akses file .env")
	}
	fileHeader, _ := c.FormFile(fileheader)
	file, _ := fileHeader.Open()
	ctx := context.Background()
	urlCloudinary := os.Getenv("CLOUDINARY_URL")
	cldService, _ := cloudinary.NewFromURL(urlCloudinary)
	response, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "Photo"})
	return response.SecureURL
}
