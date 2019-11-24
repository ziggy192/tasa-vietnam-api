package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"

	"encoding/json"

	"io/ioutil"
	"mime/multipart"

	"strconv"
)

// todo get this from config file
const ClientID = "5958ddd634f2aea"
const AccessToken = "567d191b832101282951460d490181a4ca8eb3e9"

//shit
type ProjectPost struct {
	ID        uint               `json:"id" gorm:"primary_key"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	DeletedAt *time.Time         `json:"deletedAt"`
	Title     string             `json:"title"`
	Body      string             `json:"body"`
	Subtitle  string             `json:"subittle"`
	Images    []ProjectPostImage `json:"images"`
}

type ProjectPostImage struct {
	ID            uint       `json:"id" gorm:"primary_key"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"`
	Url           string     `json:"url"`
	IsDefault     bool       `json:"isDefault"`
	ProjectPostId uint       `json:"projectPostId"`
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running...")

}
func setupDB(address string, username string, password string) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@(%s)/tasa?charset=utf8&parseTime=True&loc=Local",
		username, password, address)
	db, err := gorm.Open("mysql", connectionString)
	check(err)
	return db
}

func testPostHandler(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)

	var v map[string]interface{}
	w.Header().Set("Content-type", "application/json")
	dec.Decode(&v)
	enc.Encode(&v)

}

func insertPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)

	var p ProjectPost
	dec.Decode(&p)
	db.Create(&p)
	enc.Encode(p)
}

func getAllImagesFromPostHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	enc := json.NewEncoder(w)
	var images []ProjectPostImage
	projectPostId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	check(err)

	db.Where(&ProjectPostImage{
		ProjectPostId: uint(projectPostId),
	}).Find(&images)

	enc.Encode(images)

}

func getImageFromPostHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	enc := json.NewEncoder(w)
	imageId := mux.Vars(r)["imageId"]
	var image ProjectPostImage
	if db.First(&image, imageId).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	enc.Encode(image)
}
func insertProjectPostImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	enc := json.NewEncoder(w)
	dec := json.NewDecoder(r.Body)

	var bodyImage = ProjectPostImage{}
	dec.Decode(&bodyImage)
	postId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	check(err)

	image := ProjectPostImage{
		Url:           bodyImage.Url,
		ProjectPostId: uint(postId),
	}
	bodyImage.ProjectPostId = uint(postId)
	db.Create(&image)
	enc.Encode(image)
}

func putProjectPostImageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	imageIdStr := mux.Vars(r)["imageId"]
	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)
	var bodyImage ProjectPostImage
	dec.Decode(&bodyImage)
	var foundImage ProjectPostImage
	if db.First(&foundImage, imageIdStr).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	db.Model(&foundImage).Update("url", bodyImage.Url) // use Update instead of Updates for zero values
	enc.Encode(foundImage)
}
func deleteProjectPostImageHandler(w http.ResponseWriter, r *http.Request) {
	imageId := mux.Vars(r)["imageId"]
	if db.First(&ProjectPostImage{}, imageId).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	db.Delete(&ProjectPostImage{}, "id = ?", imageId).RecordNotFound()
	// ok
	w.WriteHeader(http.StatusOK)
}
func getAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []ProjectPost
	db.Preload("Images").Find(&posts)
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(posts)

}

func putPostHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)

	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	postId := vars["id"]

	var bodyPost ProjectPost
	dec.Decode(&bodyPost)

	var foundPost ProjectPost
	db.First(&foundPost, postId)
	if db.First(&foundPost, postId).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	db.Model(&foundPost).Updates(ProjectPost{
		Title:    bodyPost.Title,
		Body:     bodyPost.Body,
		Subtitle: bodyPost.Subtitle,
	})

	enc.Encode(foundPost)
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId := vars["id"]
	var post ProjectPost
	db.First(&post, postId)
	if db.First(&post, postId).RecordNotFound() {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	db.Delete(&post)
}
func getPostByIdHandler(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)
	var postId = vars["id"]

	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")

	var post ProjectPost
	// var images []ProjectPostImage
	//db.First(&post,postId)
	db.Preload("Images").First(&post, postId)
	if post.ID != 0 {

		enc.Encode(post)
	} else {
		w.WriteHeader(http.StatusNotFound) // used for writing status code only
	}

}
func uploadImageHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("uploadImageHandler hit")
	r.ParseMultipartForm(10 << 20)

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error when upload file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded file %s\n", fileHeader.Filename)
	fmt.Printf("File size=%v", fileHeader.Size)

	//todo upload to imgur
}

func testUploadImgur(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("assets/57232103_572879629870418_546118032422862848_n.jpg")
	check(err)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(w.Header())
	res := uploadImageToImgur(file)
	enc := json.NewEncoder(w)
	enc.Encode(&res)

}
func uploadImageToImgur(image []byte) map[string]interface{} {
	var requestBody bytes.Buffer

	multipartWriter := multipart.NewWriter(&requestBody)

	imageFieldWriter, err := multipartWriter.CreateFormField("image")

	check(err)
	_, err = imageFieldWriter.Write(image)
	check(err)

	multipartWriter.Close()

	req, err := http.NewRequest("POST", "https://api.imgur.com/3/upload", &requestBody)
	check(err)

	req.Header.Set("Authorization", fmt.Sprint("Client-ID ", ClientID))
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken))

	response, err := http.DefaultClient.Do(req)
	check(err)

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)

	fmt.Println(result)

	return result

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
func settupFlags(dbAddress *string, username *string, password *string) {
	flag.StringVar(dbAddress, "dbAddress", "", "the database address (required)")
	flag.StringVar(username, "username", "", "username of database (required)")
	flag.StringVar(password, "password", "", "password of database (required)")
	flag.Parse()
}

var db *gorm.DB

func validateFlags(flags ...interface{}) {
	for _, val := range flags {
		if val == "" {
			err := fmt.Errorf("required arguments is not inputed")
			flag.PrintDefaults()
			panic(err)
		}
	}
}
func main() {

	var dbAddress string
	var username string
	var password string
	settupFlags(&dbAddress, &username, &password)
	validateFlags(dbAddress, username, password)
	db = setupDB(dbAddress, username, password)
	if db != nil {
		defer db.Close()
	}

	//debug
	db = db.Debug()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", pingHandler)
	router.HandleFunc("/test/post", testPostHandler).Methods("POST")
	router.HandleFunc("/test/upload", testUploadImgur).Methods("GET")
	router.HandleFunc("/images/upload", uploadImageHandler).Methods("POST")
	router.HandleFunc("/posts", insertPostHandler).Methods("POST")
	router.HandleFunc("/posts", getAllPostsHandler).Methods("GET")
	router.HandleFunc("/posts/{id:[0-9]+}", getPostByIdHandler).Methods("GET")
	router.HandleFunc("/posts/{id:[0-9]+}/images", getAllImagesFromPostHanlder).Methods("GET")
	router.HandleFunc("/posts/{id:[0-9]+}/images", insertProjectPostImageHandler).Methods("POST")
	router.HandleFunc("/posts/{id:[0-9]+}/images/{imageId:[0-9]+}", getImageFromPostHanlder).Methods("GET")
	router.HandleFunc("/posts/{id:[0-9]+}/images/{imageId:[0-9]+}", deleteProjectPostImageHandler).Methods("DELETE")
	router.HandleFunc("/posts/{id:[0-9]+}/images/{imageId:[0-9]+}", putProjectPostImageHandler).Methods("PUT")
	router.HandleFunc("/posts/{id:[0-9]+}", putPostHandler).Methods("PUT")
	router.HandleFunc("/posts/{id:[0-9]+}", deletePostHandler).Methods("DELETE")

	fmt.Println("Version ===> ", version)
	fmt.Println("Server running at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
