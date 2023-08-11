package utility

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
//	Filepath "path/filepath"
)

func DecodeBase64(input string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

func UploadFiles(w http.ResponseWriter, r *http.Request, filepath string, fileLimit int64) {
	// Parse the incoming request with a maximum file size
	err := r.ParseMultipartForm(fileLimit << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	
	// Retrieve the uploaded files
	files := r.MultipartForm.File["file"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Error opening uploaded file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Create a new file on the server to store the uploaded file
		filePath := fmt.Sprintf("%s/%s", filepath, fileHeader.Filename)
		fmt.Println(filePath)
		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, "Error creating file on server", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy the contents of the uploaded file to the newly created file
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Error copying file contents", http.StatusInternalServerError)
			return
		}
	}

}

func CreateDirIfNotExist(filepath string) error {
	// Create the upload directory if it doesn't exist
	err := os.MkdirAll(filepath, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}