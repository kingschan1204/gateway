package myhttp

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var logger = log.New(os.Stderr, "main:", log.Llongfile|log.LstdFlags)

// NewProxy takes target host and creates a reverse proxy
func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req)
	}
	proxy.ModifyResponse = modifyResponse()
	proxy.ErrorHandler = errorHandler()
	return proxy, nil
}

func modifyRequest(req *http.Request) {
	//token := req.Header.Get("token")
	//user := TokenMap[token]
	//req.Header.Set("org", user.OrgId)
	req.Header.Set("X-Real-IP", req.RemoteAddr)
}

func errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Got error while modifying response: %v \n", err)
		return
	}
}

func modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		//resp.Header.Add("Access-Control-Allow-Origin", "*")
		//return errors.New("response body is invalid")
		return nil
	}
}
