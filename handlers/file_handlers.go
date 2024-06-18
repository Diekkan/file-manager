package handlers

import (
	"encoding/json"
	"filemanager/structs"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ListFilesOnDirectory(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dir := query.Get("dir")

	if dir == "" {
		dir = "./uploads" // Default directory
	}

	files, err := os.ReadDir(dir)

	if err != nil {
		http.Error(w, "Unable to list files", http.StatusInternalServerError)
		return
	}

	var fileInfos []structs.FileInfo

	for _, f := range files {

		fileInfo, err := f.Info()
		if err != nil {
			http.Error(w, "Unable to retrieve file info", http.StatusInternalServerError)
			return
		}

		ext := strings.ToLower(filepath.Ext(f.Name()))
		fileInfos = append(fileInfos, structs.FileInfo{
			Name:      f.Name(),
			Extension: ext,
			Size:      fileInfo.Size(),
			IsDir:     fileInfo.IsDir(),
		})

		//sort alphabetically
		sort.Slice(fileInfos, func(i, j int) bool {
			return fileInfo.Name()[i] < fileInfo.Name()[j]
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fileInfos)
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	dir := query.Get("dir")
	if dir == "" {
		dir = "./uploads"
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Upload limit reached", http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	os.MkdirAll(dir, os.ModePerm)

	tempFile, err := os.Create(fmt.Sprintf("%s/%s", dir, handler.Filename))
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}

	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {

		}
	}(tempFile)

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file content", http.StatusInternalServerError)
		return
	}
	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
