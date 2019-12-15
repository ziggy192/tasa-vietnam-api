package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm"
	"github.com/rs/cors"

	"github.com/ziggy192/tasa-vietnam-api/pkg/repo"
	"github.com/ziggy192/tasa-vietnam-api/pkg/router"
)

//shit

func main() {

	router := router.NewRouter()
	fmt.Println("Server running at :8000")
	defer repo.Close()
	log.Fatal(http.ListenAndServe(":8000", cors.New(cors.Options{
		AllowedMethods:   []string{"POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: false,
		Debug:            true,
	}).Handler(router)))

}
