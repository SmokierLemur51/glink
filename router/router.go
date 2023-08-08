package router

import (
	"net/http"

	"github.com/SmokierLemur51/glink/handlers"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	return mux
}