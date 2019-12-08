package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

//RecoveryHandler handle error
func RecoveryHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//wrap serveHttp with a defered recover()
		defer func() {
			// handler recovery here
			if r := recover(); r != nil {
				log.Printf("Error:%s\n", r)
				//write error message to response
				jsonBody, error := json.Marshal(map[string]string{
					"error": r.(error).Error(),
				})
				if error != nil {
					jsonBody, _ = json.Marshal(map[string]string{
						"error": "Internal Server Error",
					})
				}

				w.Header().Set("content-type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
