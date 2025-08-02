package middlewares

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"simple-project/ctxkeys"
)



type Middleware struct{}

func (m Middleware) JsonParser(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			body, _ := CheckContentTypeAndReturnBody(w, r)

			// Parsing body
			var data map[string]interface{}

			if err := json.Unmarshal(body, &data); err != nil {
				http.Error(w, "Invalid Json", http.StatusBadRequest)
				return
			}
			
			// Save data to context

			ctx := context.WithValue(r.Context(), ctxkeys.JSONBodyKey, data)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
}

func (m Middleware) JsonToStruct(w http.ResponseWriter, r *http.Request, target interface{}) error {

	body, _ := CheckContentTypeAndReturnBody(w, r)
	return json.Unmarshal(body, target)
}

func CheckContentTypeAndReturnBody(w http.ResponseWriter, r *http.Request) ([]byte, bool) {
	// Checking Content-type
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-type must be  application/json", http.StatusUnsupportedMediaType)
		return nil, false
	}

	// Reading Body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return nil, false
	}
	return body, true
}


