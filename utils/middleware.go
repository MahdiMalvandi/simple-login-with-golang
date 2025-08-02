package utils

import (
	"fmt"
	"net/http"
	"simple-project/keys"
)

func  GetJson(r *http.Request) map[string]interface{} {

	if data, ok := r.Context().Value(keys.JSONBodyKey).(map[string]interface{}); ok {
		return data

	}
	fmt.Println("data in context:",r.Context().Value(keys.JSONBodyKey))
	return nil
}