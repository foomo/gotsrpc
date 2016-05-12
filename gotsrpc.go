package gotsrpc

import (
	"encoding/json"
	"net/http"
	"strings"
)

func GetCalledFunc(r *http.Request, endPoint string) string {
	return strings.TrimPrefix(r.URL.Path, endPoint+"/")
}

func ErrorFuncNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("method not found"))
}

func ErrorCouldNotLoadArgs(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("could not load args"))
}

func ErrorMethodNotAllowed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("you gotta POST"))
}

func LoadArgs(args []interface{}, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&args)
	if err != nil {
		return err
	}
	return nil
}

func Reply(response []interface{}, w http.ResponseWriter) {
	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("could not serialize response"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
