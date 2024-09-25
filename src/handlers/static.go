package handlers

import "net/http"

// serves http static file
func HandleStatic() http.Handler {
	return http.FileServer(http.Dir("./public/static"))
}
