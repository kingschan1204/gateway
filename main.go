package main

import (
	"fmt"
	"gateway/config"
	"gateway/myhttp"
	"github.com/fatih/color"
	"log"
	"net/http"
)

//http://patorjk.com/software/taag/#p=display&f=Bloody&t=gateway
//http://patorjk.com/software/taag/#p=display&f=3D-ASCII&t=gateway
var banner string = `
 ________  ________  _________  _______   ___       __   ________      ___    ___ 
|\   ____\|\   __  \|\___   ___\\  ___ \ |\  \     |\  \|\   __  \    |\  \  /  /|
\ \  \___|\ \  \|\  \|___ \  \_\ \   __/|\ \  \    \ \  \ \  \|\  \   \ \  \/  / /
 \ \  \  __\ \   __  \   \ \  \ \ \  \_|/_\ \  \  __\ \  \ \   __  \   \ \    / / 
  \ \  \|\  \ \  \ \  \   \ \  \ \ \  \_|\ \ \  \|\__\_\  \ \  \ \  \   \/  /  /  
   \ \_______\ \__\ \__\   \ \__\ \ \_______\ \____________\ \__\ \__\__/  / /    
    \|_______|\|__|\|__|    \|__|  \|_______|\|____________|\|__|\|__|\___/ /     
                                                                     \|___|/      
`

func init() {
	color.Green(banner)
	config.InitConfig()
	color.Green("GateWay initialized with port(s): %s", config.App.Port)
}

func main() {
	// handle all requests to your server using the proxy
	http.HandleFunc("/", myhttp.ProxyRequestHandler)
	http.HandleFunc(config.App.Login, myhttp.LoginHandle)
	http.HandleFunc(config.App.Code, myhttp.GenerateCaptchaHandler)
	//http.HandleFunc("/verify", myhttp.CaptchaVerifyHandle)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.App.Port), nil))
}
