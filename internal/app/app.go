package app

import (
	"log"
	"net/http"
	"simple-proxy-image/internal/controller/http/v1/proxyImage"

	"github.com/gorilla/mux"
)

type Http struct {
	listenAddr string
}

func NewApp(listenAddr string) *Http {
	return &Http{listenAddr: listenAddr}
}

func (s *Http) Run() {
	router := mux.NewRouter()

	log.Println("Init router")

	proxyImage.InitProxyImageHandler(router)

	log.Println("HTTP server running on port: ", s.listenAddr)

	log.Fatal(http.ListenAndServe(s.listenAddr, router))
}
