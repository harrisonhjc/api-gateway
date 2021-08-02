package service

import (
	"log"
	"net/http"
	"strings"
)


type ClusterMapping struct {
	clusterMapping map[string]*Cluster
}

func NewClusterMapping() *ClusterMapping  {
	mapping := make(map[string]*Cluster)
	return &ClusterMapping{clusterMapping:mapping}
}

func (this *ClusterMapping) AddService(clusterName, serviceName, url string) {
	log.Printf("ClusterMapping:%s:%s:%s\n", clusterName, serviceName, url)
	if _, isExist := this.clusterMapping[clusterName]; !isExist {
		services := make([]*Service, 0)
		this.clusterMapping[clusterName] = &Cluster{name: clusterName, services: services}
	}
	this.clusterMapping[clusterName].addService(serviceName, url)
}

func (this *ClusterMapping) ServeRequest(w http.ResponseWriter, req *http.Request) {
	log.Println("ClusterMapping.ServeRequest+++")
	
	endpoint := strings.Trim(req.URL.Path, "/")
	serviceName := endpoint 
	if cluster, ok := this.clusterMapping[serviceName]; ok {
		cluster.serveRequest(w, req)
	}
}
