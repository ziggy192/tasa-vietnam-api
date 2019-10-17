package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
)

//shit
type ProjectPost struct {
	gorm.Model
	Title string
	Body  string
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running...")

}
func setupDB() *gorm.DB {
	db, _ := gorm.Open("mysql", "root:12345678@/tasa?charset=utf8&parseTime=True&loc=Local")

	return db

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
	log.Fatal(http.ListenAndServe(":8000", router))

}
