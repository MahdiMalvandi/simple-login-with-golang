package utils

import (
	"fmt"
	"net/http"
	"simple-project/ctxkeys"
)

func  GetJson(r *http.Request) map[string]interface{} {

	if data, ok := r.Context().Value(ctxkeys.JSONBodyKey).(map[string]interface{}); ok {
		return data

	}
	fmt.Println("data in context:",r.Context().Value(ctxkeys.JSONBodyKey))
	return nil
}