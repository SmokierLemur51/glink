package handlers

import (
	"net/http"

	"github.com/SmokierLemur51/glink/models"
	"github.com/SmokierLemur51/glink/utils"
)


// Check valid path


// func CheckValidPath(w http.ResponseWriter, r *http.Request, page models.PageData) int {
// 	if r.URL.Path != page.URL || r.Method != http.MethodGet {
// 		http.NotFound(w, r)
// 		return 
// 	} 
// 	return w.WriteHeader(http.StatusOK)
// }


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		URL: "/",
		Page: "index.html",
		Title: "glink",
	}

	if r.URL.Path != data.URL || r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)

	// CheckValidPath(w, r, data)
	err := utils.RenderTemplate(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func AboutGlinkHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		URL: "/about-glink",
		Page: "about.html",
		Title: "glink",
	}
	if r.URL.Path != data.URL || r.Method != http.MethoGet {
		http.NotFound(w, r)
		return 
	}

	err := utils.RenderTemplate(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

