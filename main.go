package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"

	"github.com/ziggy192/tasa-vietnam-api/pkg/router"
)

//shit

func main() {

	router := router.NewRouter()
	fmt.Println("Version ===> ", version)
	fmt.Println("Server running at :8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
