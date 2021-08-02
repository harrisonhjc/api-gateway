package service

import (
	"log"
	"net/http"
	"api-gateway/proxy"
)

type Service struct {
	name          string `json:"name"`
	url           string `json:"url"`
	connectionNum int32 `json:"connection_num"`
	httpProxy     *proxy.HttpProxy
}

func (this *Service) serveRequest(w http.ResponseWriter, req *http.Request) {
	log.Println("service.serveRequest +++")

	this.connectionNum++
	this.httpProxy.ServeRequest(w, req, this.url)
	this.connectionNum--
}
