package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func CheckIfPath(w http.ResponseWriter, r *http.Request, path string) {
	if r.URL.Path != path {
		fmt.Println("error: " + r.URL.Path, path)
		http.NotFound(w, r)
		return
	}
}

func IsValidHTTPMethod(method string, acceptedMethod string, w http.ResponseWriter) bool {
	if method == acceptedMethod {
		w.Header().Set("Allow", acceptedMethod)
		return true
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return false
}

var (
	EmplyFuncMap = template.FuncMap{}
	EmplyStruct = struct{}{}
)