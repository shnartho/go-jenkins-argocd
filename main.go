package main

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var (
	uploadDir = "database"
)

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/upload", uploadHandler)

	port := ":8091"
	fmt.Printf("Server is listening to port %s...", port)
	http.ListenAndServe(port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	images, err := getUploadImagesAsBase64()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title  string
		Images []string
	}{
		Title:  "Go-Jenkins",
		Images: images,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		dst, err := os.Create(filepath.Join(uploadDir, handler.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func getUploadImagesAsBase64() ([]string, error) {
	files, err := filepath.Glob(filepath.Join(uploadDir, "*"))
	if err != nil {
		return nil, err
	}

	var images []string
	for _, file := range files {
		base64Image, err := encodeImageToBase64(file)
		if err != nil {
			return nil, err
		}
		images = append(images, base64Image)
	}
	return images, nil
}

func encodeImageToBase64(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return "", err
	}

	size := fi.Size()
	buffer := make([]byte, size)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	base64Image := base64.StdEncoding.EncodeToString(buffer)
	return base64Image, nil
}
