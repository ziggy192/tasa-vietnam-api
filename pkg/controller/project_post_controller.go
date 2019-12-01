package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ziggy192/tasa-vietnam-api/pkg/model"
)

func InsertPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)

	var p model.ProjectPost
	dec.Decode(&p)
	db.Create(&p)
	enc.Encode(p)
}

func GetAllImagesFromPostHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	enc := json.NewEncoder(w)
	var images []model.ProjectPostImage
	projectPostId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	check(err)

	db.Where(&model.ProjectPostImage{
		ProjectPostId: uint(projectPostId),
	}).Find(&images)

	enc.Encode(images)

}

func GetImageFromPostHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	enc := json.NewEncoder(w)
	imageId := mux.Vars(r)["imageId"]
	var image model.ProjectPostImage
	if db.First(&image, imageId).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	enc.Encode(image)
}
func InsertProjectPostImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	enc := json.NewEncoder(w)
	dec := json.NewDecoder(r.Body)

	var bodyImage = model.ProjectPostImage{}
	dec.Decode(&bodyImage)
	postId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	check(err)

	image := model.ProjectPostImage{
		Url:           bodyImage.Url,
		ProjectPostId: uint(postId),
	}
	bodyImage.ProjectPostId = uint(postId)
	db.Create(&image)
	enc.Encode(image)
}

func PutProjectPostImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	imageIdStr := mux.Vars(r)["imageId"]
	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)
	var bodyImage model.ProjectPostImage
	dec.Decode(&bodyImage)
	var foundImage model.ProjectPostImage
	if db.First(&foundImage, imageIdStr).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	db.Model(&foundImage).Update("url", bodyImage.Url) // use Update instead of Updates for zero values
	enc.Encode(foundImage)
}
func DeleteProjectPostImageHandler(w http.ResponseWriter, r *http.Request) {
	imageId := mux.Vars(r)["imageId"]
	if db.First(&model.ProjectPostImage{}, imageId).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	db.Delete(&model.ProjectPostImage{}, "id = ?", imageId).RecordNotFound()
	// ok
	w.WriteHeader(http.StatusOK)
}
func GetAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []model.ProjectPost
	db.Preload("Images").Find(&posts)
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(posts)

}

func PutPostHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)

	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	postId := vars["id"]

	var bodyPost model.ProjectPost
	dec.Decode(&bodyPost)

	var foundPost model.ProjectPost
	db.First(&foundPost, postId)
	if db.First(&foundPost, postId).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	db.Model(&foundPost).Updates(model.ProjectPost{
		Title:    bodyPost.Title,
		Body:     bodyPost.Body,
		Subtitle: bodyPost.Subtitle,
	})

	enc.Encode(foundPost)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["id"]
	var post model.ProjectPost
	db.First(&post, postId)
	if db.First(&post, postId).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	db.Delete(&post)
}
func GetPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var postId = vars["id"]

	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")

	var post model.ProjectPost
	// var images []model.ProjectPostImage
	//db.First(&post,postId)
	db.Preload("Images").First(&post, postId)
	if post.ID != 0 {

		enc.Encode(post)
	} else {
		w.WriteHeader(http.StatusNotFound) // used for writing status code only
	}

}
