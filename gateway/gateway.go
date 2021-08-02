package gateway

import (
	"net/http"
	"api-gateway/filter"
	"api-gateway/service"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
	
	"api-gateway/config"
)

type GatewayStarter struct {
	filters []filter.IFilter
	clusterMapping *service.ClusterMapping
}

func NewGatewayStarter() *GatewayStarter {
	log.Println("GatewayStarter+++")

	starter := &GatewayStarter{
		filters: filter.HttpFilters,
		clusterMapping:service.NewClusterMapping(),
	}
	return starter
}

func (this *GatewayStarter) dispatch(w http.ResponseWriter, req *http.Request) {
	log.Printf("dispatch:%s\n", req.URL)
	this.clusterMapping.ServeRequest(w, req)
}

func (this *GatewayStarter) filter(w http.ResponseWriter, req *http.Request) bool {
	var isPass bool = true
	for _, filter := range this.filters {
		passed := filter.Filter(req)
		if isPass {
			isPass = passed
		}
	}
	return isPass
}

func (this *GatewayStarter) gateway(w http.ResponseWriter, req *http.Request) {
	isPassed := this.filter(w, req)
	if isPassed {
		this.dispatch(w, req)
	}
}

func (this *GatewayStarter) parseSetting() {
	log.Println("parseSetting++")
	
		
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal("-----------------parse config.json failed !")
		os.Exit(0)
	}
	bytes, _ := ioutil.ReadAll(file)
	vo := &config.ClusterVoList{}
	json.Unmarshal(bytes, vo)
	
	for _,cluster := range vo.Clusters {
		for _, service := range cluster.Services {
			this.clusterMapping.AddService(cluster.Name, service.Name, service.Domain)
			
		}
	}
}

func (this *GatewayStarter) Start() {
	log.Println("Start ++++++++++++")
	this.parseSetting()
	http.HandleFunc("/", this.gateway)
	http.ListenAndServe(":8080", nil)
}
