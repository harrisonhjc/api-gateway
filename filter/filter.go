package filter

import (
	"net/http"
	"log"
)

type IFilter interface {
	Filter(req *http.Request) bool
}

var HttpFilters []IFilter = []IFilter{
	LogFilter{},
}


type LogFilter struct {

}

func (this LogFilter) Filter(req *http.Request) bool {
	
	log.Println(req.Method)
	log.Println(req.URL)
	log.Println(req.Header)
	log.Println(req.Host)
	log.Println(req.RemoteAddr)
	log.Println(req.RequestURI)
	return true
}