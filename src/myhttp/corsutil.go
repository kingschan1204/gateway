package myhttp

import (
	"net/http"
)

//跨域处理
func cors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, token")
	//(*w).Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Expose-Headers", "*")
	(*w).Header().Set("Access-Control-Max-Age", "3600")
	if r.Method == "OPTIONS" {
		(*w).WriteHeader(http.StatusOK)
		return
	}
}
