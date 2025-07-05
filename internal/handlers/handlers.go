package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Read error", http.StatusInternalServerError)
		return
	}

	result := service.Convert(string(content))

	outputFilename := time.Now().UTC().Format("20060102150405") + filepath.Ext(header.Filename)
	os.WriteFile(outputFilename, []byte(result), 0644)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}
