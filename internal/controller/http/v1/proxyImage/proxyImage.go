package proxyImage

import (
	"log"
	"net/http"
	"simple-proxy-images/internal/utils"

	"github.com/gorilla/mux"
)

func InitProxyImageHandler(router *mux.Router) {
	log.Println("Init proxy image handler")

	router.HandleFunc("/proxy", proxyImage).Methods(http.MethodGet)
}

func proxyImage(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	referer := r.URL.Query().Get("referer")
	authority := r.URL.Query().Get("authority")

	image := utils.GetProxyImage(url, referer, authority)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(image)
}
