package router

import (
	"filemanager/handlers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/list-files", handlers.ListFilesOnDirectory).Methods("GET")
	router.HandleFunc("/upload-file", handlers.UploadFile).Methods("POST")
	return router
}
