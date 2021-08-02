package proxy

import (
	"net/http"
	"net/url"
	"log"
	"io"
	"io/ioutil"
)

type HttpProxy struct {
	transport *http.Transport
}

func NewHttpProxy(rawurl string) *HttpProxy {
	_, err := url.Parse(rawurl)
	if err != nil {
		log.Println(rawurl, "NewHttpProxy failed!")
		return nil
	}
	transport := &http.Transport{
		//Proxy: http.ProxyURL(url),
	}
	proxy := &HttpProxy{transport: transport}
	return proxy
}

func (this *HttpProxy) ServeRequest(w http.ResponseWriter, req *http.Request, rawurl string) {

	log.Println("HttpProxy.ServeRequest +++")

	client := &http.Client{
		Transport: this.transport,
	}
	for _, cookie := range req.Cookies() {
		req.AddCookie(cookie)
	}
	rawurl = rawurl + req.RequestURI
	req.RequestURI = ""

	log.Printf("rawurl=%s\n", rawurl)
	
	u, _ := url.Parse(rawurl)
	req.URL = u
	req.Host = u.Host
	for k, _ := range req.Header {
		if k != "Cookie" {
			req.Header.Del(k)
		}
	}
	log.Println(req)
	
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	io.Copy(w, resp.Body)
	resp.Body.Close()
}

func (this *HttpProxy) ProxyRequest(req *http.Request) []byte {
	log.Println("HttpProxy.ProxyRequest")
	
	client := &http.Client{
		Transport: this.transport,
	}
	resp, _ := client.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return data

}
