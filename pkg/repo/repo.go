package repo

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
)

func setupDB(address string, username string, password string) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@(%s)/tasa?charset=utf8&parseTime=True&loc=Local",
		username, password, address)
	DB, err := gorm.Open("mysql", connectionString)
	check(err)
	return DB
}

func settupFlags(dbAddress *string, username *string, password *string) {
	flag.StringVar(dbAddress, "dbAddress", "", "the database address (required)")
	flag.StringVar(username, "username", "", "username of database (required)")
	flag.StringVar(password, "password", "", "password of database (required)")
	flag.Parse()
}

func validateFlags(flags ...interface{}) {
	for _, val := range flags {
		if val == "" {
			err := fmt.Errorf("required arguments is not inputed")
			flag.PrintDefaults()
			panic(err)
		}
	}
}

var DB *gorm.DB

func init() {

	var dbAddress string
	var username string
	var password string
	settupFlags(&dbAddress, &username, &password)
	validateFlags(dbAddress, username, password)
	DB = setupDB(dbAddress, username, password)
	if DB != nil {
		defer DB.Close()
	}

	//debug
	DB = DB.Debug()
}
