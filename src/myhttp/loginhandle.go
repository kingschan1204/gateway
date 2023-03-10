package myhttp

import (
	"fmt"
	"gateway/src/config"
	"gateway/src/util"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	//"github.com/modern-go/reflect2"
	//"github.com/modern-go/concurrent"
)

func OpencodeHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json")
	Cors(&w, r)
	data := `{
			  "success": true,
			  "message": "操作成功",
			  "code": 200,
			  "data": {
				"base64": "data:image/jpg;base64,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",
				"id": "v:f83f81c17bb643fbbae1b507b5ff24e9"
			  },
			  "timestamp": 0
			}`
	io.WriteString(w, data)
}

type UserInfo struct {
	Id          string `json:"id"`
	OrgId       string `json:"orgId"`
	UserAccount string `json:"userAccount"`
	UserName    string `json:"userName"`
	CreateDate  int    `json:"createDate"`
	UserStatus  int    `json:"userStatus"`
	UserType    string `json:"userType"`
}

type httpResult struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Code      int         `json:"code"`
	Data      *resultData `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type resultData struct {
	Token string    `json:"token"`
	User  *UserInfo `json:"user"`
}

var TokenMap map[string]UserInfo

func LoginHandle(w http.ResponseWriter, r *http.Request) {
	url := config.App.Login + "/system/user/login"
	Cors(&w, r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read body err\n")
		return
	}

	// json
	//data := `{"userAccount":"admin","userPsw":"d033e22ae348aeb5660fc2140aec35850c4da997"}`
	result, err := HttpRequest(url, "POST", strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("get resp failed,err\n")
		return
	}
	//fmt.Println(string(result))
	w.Header().Set("content-type", "text/json")

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//是否成功
	success := jsoniter.Get(result, "success").ToBool()
	message := jsoniter.Get(result, "message").ToString()
	//code:=jsoniter.Get(result, "code").ToInt()
	user := jsoniter.Get(result, "data").ToString()
	var uinfo UserInfo
	if success {
		json.Unmarshal([]byte(user), &uinfo)
	}
	//token
	token := util.GetMd5(uinfo.Id)
	resultData := resultData{token, &uinfo}
	httpresult := httpResult{success, message, 200, nil, time.Now().Unix()}
	if success {
		httpresult.Data = &resultData
	}
	b, err := json.Marshal(httpresult)
	TokenMap[token] = uinfo
	io.WriteString(w, string(b))

}
