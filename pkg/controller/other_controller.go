package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// PingHandler ping
func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running...")
	log.Println("ping")
}

//PanicHandler throw dummy panic
func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic(errors.New("[Test] Error in PanicHanlder"))
}

func TestPostHandler(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)

	var v map[string]interface{}
	w.Header().Set("Content-type", "application/json")
	dec.Decode(&v)
	enc.Encode(&v)

}
