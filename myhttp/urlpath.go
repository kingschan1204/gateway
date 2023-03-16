package myhttp

import (
	"gateway/config"
	"regexp"
	"strings"
)

// Determine whether the path is a whitelist
// url ：request url
// whiteList：config file whitelist
func whiteList(url string, whiteList []string) bool {
	exists := false
	for i := 0; i < len(whiteList); i++ {
		if m, _ := regexp.MatchString(whiteList[i], url); m {
			exists = true
			break
		}
	}
	return exists
}

// isRoutePath Determine whether the path is need route
//return : 1. is route , 2. route path , 3.repace path prefix
func isRoutePath(url string, route map[string]*config.RouteInfo, depth int) (bool, *config.RouteInfo, string) {
	path := ""
	if url == "/" {
		path = "/"
	} else {
		result := strings.Split(url, "/")
		var array []string
		array = result[1:]
		for i := 0; i < depth; i++ {
			if i < len(array) && array[i] != "" {
				path += "/" + array[i]
			}
		}
	}
	v, ok := route[path]
	if ok {
		return true, v, path
	}
	return false, nil, ""
}
