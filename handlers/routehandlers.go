package handlers

import (
	"net/http"

	"github.com/SmokierLemur51/glink/models"
	"github.com/SmokierLemur51/glink/utils"
)


// Check valid path


func CheckValidPath(w http.ResponseWriter, r *http.Request, page models.PageData) int {
	if r.URL.Path != page.URL || r.Method != http.MethodGet {
		http.NotFound(w, r)
		return 
	} 
	return w.WriteHeader(http.StatusOK)
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		URL: "/",
		Page: "index.html",
		Title: "glink",
	}
	CheckValidPath(w, r, data)
	err := utils.RenderTemplate(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

