package proxyImage

import (
	"log"
	"net/http"
	"simple-proxy-image/internal/utils"

	"github.com/gorilla/mux"
)

func InitProxyImageHandler(router *mux.Router) {
	log.Println("Init proxy image handler")

	router.HandleFunc("/proxy", getProxedImage).Methods(http.MethodGet)
}

func getProxedImage(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	referer := r.URL.Query().Get("referer")
	authority := r.URL.Query().Get("authority")

	image, err := utils.FetchImage(url, referer, authority)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(image)
}
