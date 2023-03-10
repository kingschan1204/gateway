package myhttp

import (
	"gateway/src/config"
	"io"
	"net/http"
	"strings"
)

var result = `{
			  "success": false,
			  "message": "token don't exist or has expired",
			  "code": 401,
			  "data": { },
			  "timestamp": 0
			}`

// ProxyRequestHandler handles the myhttp request using proxy
//proxy *httputil.ReverseProxy
func ProxyRequestHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		token := r.Header.Get("token")
		host := ""
		cors(&w, r)

		isRoute, routePath, replacePath := isRoutePath(url, config.App.Route, config.App.RouteDepth)
		if isRoute {
			r.URL.Path = strings.Replace(r.URL.Path, replacePath, "", 1)
			host = routePath

		}
		// If it is a whitelist directly through
		if !whiteList(url, config.App.WhiteList) {
			// need token
			if _, ok := TokenMap[token]; token == "" || !ok {
				w.Header().Set("content-type", "text/json")
				io.WriteString(w, result)
				return
			}
		}
		proxy, _ := NewProxy(host)
		proxy.ServeHTTP(w, r)
	}
}
