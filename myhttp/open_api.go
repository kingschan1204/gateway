package myhttp

import (
	"fmt"
	"gateway/app"
	"gateway/plugin"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strings"
)

func GenerateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, b64s, err := plugin.CaptchaGenerate()
	body := map[string]interface{}{"data": b64s, "id": id, "code": 200}
	if err != nil {
		body = map[string]interface{}{"code": 500, "msg": err.Error()}
	}
	jsoniter.NewEncoder(w).Encode(body)
}

type loginResult struct {
	Message string     `json:"message"`
	Code    int        `json:"code"`
	Data    *LoginData `json:"data"`
}
type LoginData struct {
	Username string `json:"username"`
	Tenant   string `json:"tenant"`
}

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read body err")
		return
	}
	result, err := HttpRequest(app.Config.LoginApi, "POST", strings.NewReader(string(body)))
	if err != nil {
		fmt.Println(err)
		return
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var login loginResult
	json.Unmarshal(result, &login)

	rd := map[string]interface{}{"msg": "login error ", "code": 500}
	if login.Code == 200 {
		token, err := plugin.GenToken(login.Data.Username, login.Data.Tenant, []byte(app.Config.TokenSecret), app.Config.TokenExpire)
		if nil == err {
			rd = map[string]interface{}{"token": token, "code1": 200}
		}
	}
	json.NewEncoder(w).Encode(rd)

}
