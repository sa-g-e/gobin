package handlers

import (
    "html/template"
    "net/http"
    "path/filepath"
    "gobin/storage"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
    pasteID := r.URL.Path[len("/view/"):]

    paste, err := storage.LoadPaste(pasteID)
    if err != nil {
        http.Error(w, "Paste not found", http.StatusNotFound)
        return
    }

    tmplPath := filepath.Join("static", "view.html")
    tmpl := template.Must(template.ParseFiles(tmplPath))

    if err := tmpl.Execute(w, paste); err != nil {
        http.Error(w, "Failed to render paste", http.StatusInternalServerError)
    }
}
