package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"

	"encoding/json"
)

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
func setupDB() *gorm.DB {
	db, _ := gorm.Open("mysql", "root:12345678@/tasa?charset=utf8&parseTime=True&loc=Local")

	return db
}

func testPostHandler(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)

	var v map[string]interface{}
	dec.Decode(&v)
	enc.Encode(&v)

}

func main() {
	db := setupDB()
	//p := ProjectPost{Title: "công trình abc", Body: "adsfas fsa fas dfas df as"}
	//db.Create(&p)
	pselect := ProjectPost{}
	db.Where("id = ?", 1).First(&pselect)
	fmt.Println(pselect)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", pingHandler)
	router.HandleFunc("/testPost", testPostHandler).Methods("POST")

	fmt.Println("Server running at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
