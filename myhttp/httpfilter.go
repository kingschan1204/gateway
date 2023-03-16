package myhttp

import (
	"fmt"
	"gateway/config"
	"gateway/plugin"
	"io"
	"net/http"
	"strings"
)

var result = `{
			  "message": "token don't exist or has expired",
			  "code": 401,
			}`

// ProxyRequestHandler handles the myhttp request using proxy
//proxy *httputil.ReverseProxy
func ProxyRequestHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	token := r.Header.Get("token")
	host := ""
	cors(&w, r)

	isRoute, routePath, replacePath := isRoutePath(url, config.App.Route, config.App.RouteDepth)
	if isRoute {
		r.URL.Path = strings.Replace(r.URL.Path, replacePath, "", 1)
		// To achieve load balancing in the future
		host = routePath[0]

	}
	// If it is a whitelist directly through
	if !whiteList(url, config.App.WhiteList) {
		// need token
		tokenClaims, err := plugin.ParseToken(token, []byte(config.App.TokenSecret))
		if err != nil {
			w.Header().Set("content-type", "text/json")
			w.WriteHeader(401)
			io.WriteString(w, result)
			return
		}
		// the token is ok .so set http head data
		r.Header.Set("user", tokenClaims.Username)
		r.Header.Set("tenant", tokenClaims.Tenant)
		//r.Header.Set("X-Real-IP", r.RemoteAddr)
		r.Header.Set("Host", r.Host)
		r.Header.Set("Origin", fmt.Sprintf("%s://%s", r.URL.Scheme, r.Host))
		r.Header.Set("Referrer", fmt.Sprintf("%s://%s%s", r.URL.Scheme, r.Host, r.URL.RawPath))
		//if _, ok := TokenMap[token]; token == "" || !ok {
		//	w.Header().Set("content-type", "text/json")
		//	io.WriteString(w, result)
		//	return
		//}
	}
	proxy, _ := NewProxy(host)
	proxy.ServeHTTP(w, r)
}
