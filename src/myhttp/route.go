package myhttp

import (
	"bytes"
	"errors"
	"fmt"
	"gateway/src/config"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
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
	//proxy.ModifyResponse = modifyResponse()
	proxy.ModifyResponse = func(response *http.Response) error {
		cont, _ := ioutil.ReadAll(response.Body)
		response.Body = ioutil.NopCloser(bytes.NewReader(cont))
		return nil
	}
	proxy.ErrorHandler = errorHandler()
	return proxy, nil
}

func modifyRequest(req *http.Request) {
	token := req.Header.Get("token")
	//req.Header.Set("X-Proxy", "Simple-Reverse-Proxy")

	user := TokenMap[token]
	req.Header.Set("org", user.OrgId)
	req.Header.Set("account", user.UserAccount)
	req.Header.Set("uid", user.Id)
	req.Header.Set("type", user.UserType)
}

func errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		fmt.Printf("Got error while modifying response: %v \n", err)
		return
	}
}

func modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		resp.Header.Add("Access-Control-Allow-Origin", "*")
		return errors.New("response body is invalid")
	}
}

// ProxyRequestHandler handles the myhttp request using proxy
//proxy *httputil.ReverseProxy
func ProxyRequestHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("Origin:" + r.Header.Get("Origin"))
		result := `{
			  "success": false,
			  "message": "token为空或者已失效！",
			  "code": 401,
			  "data": { },
			  "timestamp": 0
			}`
		Cors(&w, r)
		var url = r.URL.Path
		token := r.Header.Get("token")
		host := ""
		log.Println("current request url ->", url)
		isRoute, routePath, replacePath := isRoutePath(url, config.App.Route, config.App.RouteDepth)
		if isRoute {
			fmt.Println("replace ->", replacePath)
			r.URL.Path = strings.Replace(r.URL.Path, replacePath, "", 1)
			host = routePath
			log.Println("route url ->", r.URL.Path)
			log.Println("referer ->", r.Header.Get("referer"))

		}
		// If it is a whitelist directly through
		if !whiteList(url, config.App.WhiteList) {
			// 需要token的
			if token == "" {
				w.Header().Set("content-type", "text/json")
				io.WriteString(w, result)
				return
			}
			//验证token是否存在
			_, ok := TokenMap[token]
			if !ok {
				fmt.Println("token为空或者失效！")
				w.Header().Set("content-type", "text/json")
				io.WriteString(w, result)
				return
			}
		}
		proxy, err1 := NewProxy(host)
		if err1 != nil {
			panic(err1)
		}
		proxy.ServeHTTP(w, r)
	}
}
