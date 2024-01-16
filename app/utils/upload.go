package utils

import (
	"io"
	"mime/multipart"
	"os"
)

func UploadFile(file *multipart.FileHeader, folder string) (string, error) {
	// Get the file from the request
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create a destination file
	// dst, err := os.Create("uploads/" + file.Filename)
	path := "uploads/" + folder + "/" + file.Filename
	dst, err := os.Create("public/" + path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return path, nil
}
