package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// PingHandler ping
func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running...")
	log.Println("ping")
}

func TestPostHandler(w http.ResponseWriter, r *http.Request) {

	dec := json.NewDecoder(r.Body)
	enc := json.NewEncoder(w)

	var v map[string]interface{}
	w.Header().Set("Content-type", "application/json")
	dec.Decode(&v)
	enc.Encode(&v)

}
