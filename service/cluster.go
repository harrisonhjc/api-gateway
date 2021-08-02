package service

import (
	"log"
	"net/http"
	"api-gateway/proxy"
)

type Cluster struct {
	name     string `json:"name"`
	services []*Service
}

func (this *Cluster) addService(name, url string) {
	log.Printf("Cluster:%s:%s\n", name, url)
	service := &Service{
		name: name,
		url: url,
		connectionNum:0,
		httpProxy:proxy.NewHttpProxy(url),
	}
	this.services = append(this.services, service)
}

func (this *Cluster) serveRequest(w http.ResponseWriter, req *http.Request)  {
	log.Println("cluster:serveRequest+++")
	log.Println(this.services)

	var serviceImpl *Service = this.services[0]
	for _, v := range this.services {
		log.Printf("connectionNum: %d:%d\n", serviceImpl.connectionNum, v.connectionNum)
		if serviceImpl.connectionNum > v.connectionNum {
			serviceImpl = v
		}
	}
	serviceImpl.serveRequest(w, req)
}