package myhttp

import (
	"fmt"
	"gateway/config"
	"gateway/util"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	//"github.com/modern-go/reflect2"
	//"github.com/modern-go/concurrent"
)

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

func init() {
	TokenMap = make(map[string]UserInfo)
}

func LoginHandle1(w http.ResponseWriter, r *http.Request) {
	url := config.App.Login + "/system/user/login"
	cors(&w, r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read body err")
		return
	}

	// json
	//data := `{"userAccount":"admin","userPsw":"d033e22ae348aeb5660fc2140aec35850c4da997"}`
	result, err := HttpRequest(url, "POST", strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("get resp failed,err")
		return
	}
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
