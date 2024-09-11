package handlers

import (
	"log"
    "net/http"
    "time"
    "gobin/storage"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    content := r.FormValue("content")
    expireStr := r.FormValue("expire")

    expire, err := time.ParseDuration(expireStr)
    if err != nil {
        expire = 0 
    }

    id, err := storage.SavePaste(content, expire)
    if err != nil {
        http.Error(w, "Failed to save paste", http.StatusInternalServerError)
		//debugging
		log.Println(err)
        return
    }

    http.Redirect(w, r, "/view/"+id, http.StatusSeeOther)
}
