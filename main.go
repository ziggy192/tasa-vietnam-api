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
	Tags      string
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running...")

}
func setupDB(address string) *gorm.DB {
	connectionString := fmt.Sprintf("root:12345678@(%s)/tasa?charset=utf8&parseTime=True&loc=Local", address)
	db, _ := gorm.Open("mysql", connectionString)

	return db
}

func testPostHandler(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)

	var v map[string]interface{}
	dec.Decode(&v)
	enc.Encode(&v)

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
func settupFlags(dbAddress *string) {
	flag.StringVar(dbAddress, "address", "localhost:3306", "the database address")
	flag.Parse()
}
func main() {

	var dbAddress string
	settupFlags(&dbAddress)
	db := setupDB(dbAddress)

	//p := ProjectPost{Title: "công trình abc", Body: "adsfas fsa fas dfas df as"}
	//db.Create(&p)
	fmt.Println(("trying query in project_posts..."))
	pselect := ProjectPost{}
	db.Where("id = ?", 1).First(&pselect)
	fmt.Println(pselect)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", pingHandler)
	router.HandleFunc("/test/post", testPostHandler).Methods("POST")
	router.HandleFunc("/test/upload", testUploadImgur).Methods("GET")
	router.HandleFunc("/images/upload", uploadImageHandler).Methods("POST")

	fmt.Println("Version ===> ", version)
	fmt.Println("Server running at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
