package controller

import "net/http"

import "encoding/json"

//GetSectionsHandler get hard-coded sections
func GetSectionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	enc := json.NewEncoder(w)
	sections := []string{"project", "inspiration", "style"}
	enc.Encode(sections)
}
