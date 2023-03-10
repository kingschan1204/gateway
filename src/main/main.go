package main

import (
	"fmt"
	"gateway/src/config"
	"gateway/src/myhttp"
	"log"
	"net/http"
)

//bloody
//http://patorjk.com/software/taag/#p=display&f=Bloody&t=gateway
var banner string = `

  ▄████  ▄▄▄     ▄▄▄█████▓▓█████  █     █░ ▄▄▄     ▓██   ██▓
 ██▒ ▀█▒▒████▄   ▓  ██▒ ▓▒▓█   ▀ ▓█░ █ ░█░▒████▄    ▒██  ██▒
▒██░▄▄▄░▒██  ▀█▄ ▒ ▓██░ ▒░▒███   ▒█░ █ ░█ ▒██  ▀█▄   ▒██ ██░
░▓█  ██▓░██▄▄▄▄██░ ▓██▓ ░ ▒▓█  ▄ ░█░ █ ░█ ░██▄▄▄▄██  ░ ▐██▓░
░▒▓███▀▒ ▓█   ▓██▒ ▒██▒ ░ ░▒████▒░░██▒██▓  ▓█   ▓██▒ ░ ██▒▓░
 ░▒   ▒  ▒▒   ▓▒█░ ▒ ░░   ░░ ▒░ ░░ ▓░▒ ▒   ▒▒   ▓▒█░  ██▒▒▒ 
  ░   ░   ▒   ▒▒ ░   ░     ░ ░  ░  ▒ ░ ░    ▒   ▒▒ ░▓██ ░▒░ 
░ ░   ░   ░   ▒    ░         ░     ░   ░    ░   ▒   ▒ ▒ ░░  
      ░       ░  ░           ░  ░    ░          ░  ░░ ░     
                                                    ░ ░
`

func init() {
	fmt.Println(banner)
	config.InitConfig()
	//login.TokenMap = make(map[string]login.UserInfo)
}

func main() {
	// handle all requests to your server using the proxy
	http.HandleFunc("/", myhttp.ProxyRequestHandler())
	http.HandleFunc(config.App.Login, myhttp.LoginHandle)
	http.HandleFunc(config.App.Code, myhttp.OpencodeHandle)
	log.Fatal(http.ListenAndServe(":"+config.App.Port, nil))
}
