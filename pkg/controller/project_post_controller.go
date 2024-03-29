package controller

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ziggy192/tasa-vietnam-api/pkg/dto/response"
	"github.com/ziggy192/tasa-vietnam-api/pkg/model"
)

//todo take this from config file
const DEFAULT_LIMIT = 100
const DEFAULT_OFFSET = 0

func InsertPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)

	var p model.ProjectPost
	dec.Decode(&p)
	if err := db.Create(&p).Error; err != nil {
		panic(err)
	}
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
	check(db.Delete(&model.ProjectPostImage{}, "id = ?", imageId).Error)
	// ok
	w.WriteHeader(http.StatusOK)
}

//return nil if error
func getQueryParamInt(params url.Values, key string) (int, error) {
	var ret int
	str := params.Get(key)
	valueInt64, err := strconv.ParseInt(str, 10, 32)
	if err == nil {
		ret = int(valueInt64)
	}
	return ret, err

}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {

	limit, err := getQueryParamInt(r.URL.Query(), "limit")
	if err != nil {
		limit = DEFAULT_LIMIT
	}
	offset, err := getQueryParamInt(r.URL.Query(), "offset")

	if err != nil {
		offset = DEFAULT_OFFSET
	}

	section := r.URL.Query().Get("section")
	var posts []model.ProjectPost
	var total int
	// query all if section==""
	db.Preload("Images").Where(&model.ProjectPost{
		Section: section,
	}).Limit(limit).Offset(offset).Find(&posts).Count(&total)

	responseDTO := response.BaseListResponse{
		Total:  total,
		Limit:  limit,
		Offset: offset,
		Data:   posts,
	}

	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(responseDTO)

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
	err := db.Model(&foundPost).Updates(model.ProjectPost{
		Title:    bodyPost.Title,
		Body:     bodyPost.Body,
		Subtitle: bodyPost.Subtitle,
		Section:  bodyPost.Section,
	}).Error
	check(err)

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
	check(db.Delete(&post).Error)
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
