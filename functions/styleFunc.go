package functions

import (
	"net/http"
	"os"
	"strings"
)

// StyleFunc serves static file style.css from the "static" directory.
func StyleFunc(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(r.URL.Path, "/")
	File, err := os.Stat(filePath)
	if err != nil || File.IsDir() {
		HandleError(w, "Not Found!", http.StatusNotFound)
		return
	}
	http.StripPrefix("/static", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
}
