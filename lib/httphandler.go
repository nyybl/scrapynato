package lib

import (
	"encoding/json"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) ResponseSchema

/*
* Wraps a custom Handler function into an http.HandlerFunc.

Parameters:
f : Handler<func(w http.ResponseWriter, r *http.Request) ResponseSchema>

Returns:
http.handlerFunc
*/
func HandleHttp(f Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		res := f(w, r)
		if err := writeJSON(res, w); err != nil {
			writeBytes(err.Error(), w)
		}
	}
}

func writeJSON(r ResponseSchema, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	return json.NewEncoder(w).Encode(r)
}

/*
* Write bytes to the client in case writeJSON fails
*/
func writeBytes(c string, w http.ResponseWriter) {
	w.Write([]byte(c))
}