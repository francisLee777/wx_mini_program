package util

import (
	"encoding/json"
	"net/http"
)

func BindJson(r *http.Request, ptr interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(ptr)
}
