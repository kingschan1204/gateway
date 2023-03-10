package myhttp

import (
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

// Determine whether the path is need route
//return is route , route path , repace path prefix
func isRoutePath(url string, route map[string]string, depth int) (bool, string, string) {
	result := strings.Split(url, "/")
	var array []string
	array = result[1:]
	path := ""
	for i := 0; i < depth; i++ {
		if i < len(array) && array[i] != "" {
			path += "/" + array[i]
		}
	}
	v, ok := route[path]
	if ok {
		return true, v, path
	}
	return false, "", ""
}
