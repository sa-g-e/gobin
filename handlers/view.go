package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/sa-g-e/gobin/storage"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	pasteID := r.URL.Path[len("/view/"):]
	var paste *[]storage.Paste
	var err error
	if pasteID == "" {
		paste, err = storage.LoadAllPaste()
	} else {
		paste, err = storage.LoadPaste(pasteID)
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("Paste not found %v", err), http.StatusNotFound)
		return
	}

	tmplPath := filepath.Join("static", "view.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))

	if err := tmpl.Execute(w, paste); err != nil {
		http.Error(w, fmt.Sprintf("Failed to render %v", err), http.StatusInternalServerError)
		return
	}
}
