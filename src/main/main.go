package main

import (
	"fmt"
	"gateway/src/config"
	"gateway/src/myhttp"
	"log"
	"net/http"
)

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
}

func main() {
	// handle all requests to your server using the proxy
	http.HandleFunc("/", myhttp.ProxyRequestHandler())
	http.HandleFunc(config.App.Login, myhttp.LoginHandle)
	http.HandleFunc(config.App.Code, myhttp.GenerateCaptchaHandler)
	//http.HandleFunc("/verify", myhttp.CaptchaVerifyHandle)
	log.Fatal(http.ListenAndServe(":"+config.App.Port, nil))
}
