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
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Body      string
	Subtitle string 
	Images 	[]ProjectPostImage
}

type ProjectPostImage struct {
	ID 			uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	url 	string
	IsDefault	bool
	ProjectPostId uint
	
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


func getAllImagesFromPostHanlder(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	enc := json.NewEncoder(w)
	var images []ProjectPostImage
	projectPostId, err := strconv.ParseUint(mux.Vars(r)["postId"],10,32)
	check(err)

	db.Where(&ProjectPostImage{			
		ProjectPostId: uint(projectPostId),
	}).Find(&images)

	enc.Encode(images)



}
func getAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []ProjectPost
	db.Preload("Images").Find(&posts)
	enc := json.NewEncoder(w)
	w.Header().Set("content-type", "application/json")
	enc.Encode(posts)

}

func getPostByIdHandler(w http.ResponseWriter, r *http.Request){
	var vars = mux.Vars(r)
	var postId = vars["postId"]
	
	enc := json.NewEncoder(w)
	w.Header().Set("content-type","application/json")
	var post ProjectPost
	// var images []ProjectPostImage
	//db.First(&post,postId)
	db.Preload("Images").First(&post,postId)
	enc.Encode(post)

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

func validateFlags(flags ...interface{}){
	for _, val := range flags {
		if val == "" {
			err:= fmt.Errorf("required arguments is not inputed")
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

	//p := ProjectPost{Title: "công trình abc", Body: "adsfas fsa fas dfas df as"}
	//db.Create(&p)
	// fmt.Println(("trying query in project_posts..."))
	// pselect := ProjectPost{}
	// db.Where("id = ?", 1).First(&pselect)
	// fmt.Println(pselect)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", pingHandler)
	router.HandleFunc("/test/post", testPostHandler).Methods("POST")
	router.HandleFunc("/test/upload", testUploadImgur).Methods("GET")
	router.HandleFunc("/images/upload", uploadImageHandler).Methods("POST")
	router.HandleFunc("/posts", insertPostHandler).Methods("POST")
	router.HandleFunc("/posts", getAllPostsHandler).Methods("GET")
	router.HandleFunc("/posts/{postId:[0-9]+}",getPostByIdHandler).Methods("GET")
	router.HandleFunc("/posts/{postId:[0-9]+}/images", getAllImagesFromPostHanlder).Methods("GET")

	fmt.Println("Version ===> ", version)
	fmt.Println("Server running at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
