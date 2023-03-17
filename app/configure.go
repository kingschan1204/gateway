package app

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type appConfig struct {
	Port        string          `yaml:"port"`        // bind port
	Code        string          `yaml:"codeUrl"`     // get validate code url
	Login       string          `yaml:"loginUrl"`    // login url
	LoginApi    string          `yaml:"loginApi"`    // login api http url
	TokenSecret string          `yaml:"tokenSecret"` // jwt token secret
	TokenExpire string          `yaml:"tokenExpire"` // jwt token Expire time .the unit is Second
	Svc         map[string]*svc `yaml:"service"`     // services key:service id
	RouteDepth  int             `yaml:"routeDepth"`  // url prefix route depth
	HostRoute   []*hostRouter   `yaml:"hostRouter"`  // host router key : host  value : service id
	PrefixRoute []*PrefixRouter `yaml:"prefixRoute"` // prefix router - > key : url path  value : service id
	WhiteList   []string        `yaml:"whiteList"`   // whitelist

}

type Yaml struct {
	App appConfig `yaml:"gateway"`
}

//define svc
type svc struct {
	Urls []string `yaml:"urls"` // service url
}

type hostRouter struct {
	Host    string `yaml:"host"`
	Service string `yaml:"service"`
}

type PrefixRouter struct {
	Path        string `yaml:"path"`
	StripPrefix bool   `yaml:"stripPrefix"`
	Service     string `yaml:"service"`
}

////////////////////////////////////////////////////////

var Config appConfig

func InitConfig() {
	//获取当前目录
	//fmt.Println(os.Getwd())
	filename := "./gateway.yaml"
	y := new(Yaml)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read gateway.yaml file error %v\n", err)
	}
	err = yaml.Unmarshal(yamlFile, y)
	if err != nil {
		log.Fatalf("yaml 解码失败: %v\n", err)
	}
	Config = y.App

	// setting router mapping
	// host mapping
	// clean map
	HostRouterMapping = make(map[string]string)
	for i := 0; i < len(Config.HostRoute); i++ {
		key := Config.HostRoute[i].Host
		HostRouterMapping[key] = Config.HostRoute[i].Service
	}
	// setting prefix router
	PrefixRouterMapping = make(map[string]*PrefixRouter)
	for i := 0; i < len(Config.PrefixRoute); i++ {
		key := Config.PrefixRoute[i].Path
		PrefixRouterMapping[key] = Config.PrefixRoute[i]
	}
	fmt.Println("init ...", HostRouterMapping)
}
