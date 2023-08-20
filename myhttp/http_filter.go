package myhttp

import (
	"fmt"
	"gateway/app"
	"gateway/plugin"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var logHttp = log.New(os.Stdout, "http:", log.Llongfile|log.LstdFlags)
var result = `{
			  "message": "token don't exist or has expired",
			  "code": 401,
			}`

// ProxyRequestHandler handles the myhttp request using proxy
//proxy *httputil.ReverseProxy
func ProxyRequestHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	token := r.Header.Get("token")
	routeHost := ""
	cors(&w, r)

	// host router
	svcId, exists := app.HostRouterMapping[r.Host]
	if exists {
		svc, ok := app.Config.Svc[svcId]
		if ok {
			r.Header.Set("router", "host")
			proxy, _ := NewProxy(svc.Urls[0])
			proxy.ServeHTTP(w, r)
			return
		}

	}

	isRoute, prefixRoute, replacePath := isRoutePath(url, app.PrefixRouterMapping, app.Config.RouteDepth)
	if isRoute {
		r.Header.Set("prefixRoute-path", replacePath)
		if !prefixRoute.StripPrefix {
			r.URL.Path = strings.Replace(r.URL.Path, replacePath, "", 1)
		}
		// To achieve load balancing in the future
		svc, ok := app.Config.Svc[prefixRoute.Service]
		if ok {
			routeHost = svc.Urls[0]
		}
		r.Header.Set("router", "prefix")

	} else {
		w.WriteHeader(404)
	}
	// If it is a whitelist directly through
	if !whiteList(url, app.Config.WhiteList) {
		// need token
		tokenClaims, err := plugin.ParseToken(token, []byte(app.Config.TokenSecret))
		if err != nil {
			w.Header().Set("content-type", "text/json")
			//w.WriteHeader(401)
			logProxy.Println("401:", r.URL)
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
	proxy, _ := NewProxy(routeHost)
	proxy.ServeHTTP(w, r)
}
