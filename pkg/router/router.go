package router

import (
	"github.com/gorilla/mux"

	"github.com/ziggy192/tasa-vietnam-api/pkg/controller"
)

//create new router from controllers
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", controller.PingHandler)
	router.HandleFunc("/test/post", controller.TestPostHandler).Methods("POST")
	router.HandleFunc("/test/upload", controller.TestUploadImgur).Methods("GET")
	router.HandleFunc("/images/upload", controller.UploadImageHandler).Methods("POST")
	router.HandleFunc("/posts", controller.InsertPostHandler).Methods("POST")
	router.HandleFunc("/posts", controller.GetAllPostsHandler).Methods("GET")
	router.HandleFunc("/posts/{id:[0-9]+}", controller.GetPostByIdHandler).Methods("GET")
	router.HandleFunc("/posts/{id:[0-9]+}/images", controller.GetAllImagesFromPostHanlder).Methods("GET")
	router.HandleFunc("/posts/{id:[0-9]+}/images", controller.InsertProjectPostImageHandler).Methods("POST")
	router.HandleFunc("/images/{imageId:[0-9]+}", controller.GetImageFromPostHanlder).Methods("GET")
	router.HandleFunc("/images/{imageId:[0-9]+}", controller.DeleteProjectPostImageHandler).Methods("DELETE")
	router.HandleFunc("/images/{imageId:[0-9]+}", controller.PutProjectPostImageHandler).Methods("PUT")
	router.HandleFunc("/posts/{id:[0-9]+}", controller.PutPostHandler).Methods("PUT")
	router.HandleFunc("/posts/{id:[0-9]+}", controller.DeletePostHandler).Methods("DELETE")
	return router
}
