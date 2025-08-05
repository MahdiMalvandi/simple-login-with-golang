package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"simple-project/keys"
	"simple-project/utils"
)

type Middleware struct{}

func (m Middleware) JsonParser(next http.Handler, logger *utils.Logger) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			body, _ := CheckContentTypeAndReturnBody(w, r)
			host, _, _ := net.SplitHostPort(r.RemoteAddr)
			logger.Log(fmt.Sprintf("INFO:Middleware->JsonParser: Content-Type was True for route: %s , Method: %s, User's IP: %s", r.URL.Path, r.Method, host))
			// Parsing body
			var data map[string]interface{}

			if err := json.Unmarshal(body, &data); err != nil {
				http.Error(w, "Invalid Json", http.StatusBadRequest)
				logger.Log(fmt.Sprintf("ERROR:Middleware->JsonParser: Invalid Json route: %s , Method: %s, User's IP: %s", r.URL.Path, r.Method, host))

				return
			}

			// Save data to context
			ctx := context.WithValue(r.Context(), keys.JSONBodyKey, data)
			logger.Log(fmt.Sprintf("INFO:Middleware->JsonParser: Parsing Json Was Successfull for route: %s , Method: %s, User's IP: %s", r.URL.Path, r.Method, host))

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
