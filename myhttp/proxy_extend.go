package myhttp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var logProxy = log.New(os.Stderr, "proxy:", log.Llongfile|log.LstdFlags)

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

func modifyRequest(r *http.Request) {
	r.Header.Del("token")
}

func errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("Got error while modifying response: %v \n", err)
		return
	}
}

func modifyResponse() func(*http.Response) error {
	return func(res *http.Response) error {
		//resp.Header.Add("Access-Control-Allow-Origin", "*")
		//return errors.New("response body is invalid")
		contentType := res.Header.Get("Content-Type")
		encoding := res.Header.Get("Content-Encoding")

		// contentType: application/json
		// contentType: application/javascript
		// contentType: image/png
		// contentType: text/css

		fmt.Println("contentType ->", contentType)
		fmt.Println("encoding ->", encoding)

		switch encoding {
		case "gzip":
			routePrefix := res.Request.Header.Get("route-path")
			host := res.Request.Host
			color.Red("modify response : %s", res.Request.URL)
			color.Red("Host : %s", host)
			color.Red("URL : %s", res.Request.URL)
			color.Red("route-path : %s", routePrefix)
			color.Red("RequestURI : %s", res.Request.RequestURI)
			fmt.Println("")
			reader, err := gzip.NewReader(res.Body)

			if err != nil {
				return errors.WithStack(err)
			}

			defer reader.Close()

			body, err := ioutil.ReadAll(reader)
			html := string(body)

			regex := regexp.MustCompile(`https?://\S+`)
			urls := regex.FindAllString(html, -1)
			for i := 0; i < len(urls); i++ {
				old := urls[i]
				new := strings.ReplaceAll(old, res.Request.Host, host+routePrefix)
				html = strings.ReplaceAll(html, old, new)

			}
			if err != nil {
				return errors.WithStack(err)
			}
			newBody := html

			var b bytes.Buffer
			gz := gzip.NewWriter(&b)

			if _, err := gz.Write([]byte(newBody)); err != nil {
				return errors.WithStack(err)
			}

			if err := gz.Close(); err != nil {
				return errors.WithStack(err)
			}

			bin := b.Bytes()
			res.Header.Set("Content-Length", fmt.Sprint(len(bin)))
			res.Body = io.NopCloser(bytes.NewReader(bin))

		}
		return nil
	}
}
