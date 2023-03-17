package app

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type appConfig struct {
	// bind port
	Port string `yaml:"port"`
	// get validate code url
	Code string `yaml:"codeUrl"`
	// login url
	Login string `yaml:"loginUrl"`
	// login api http url
	LoginApi string `yaml:"loginApi"`
	// jwt token secret
	TokenSecret string `yaml:"tokenSecret"`
	// jwt token Expire time .the unit is Second
	TokenExpire string `yaml:"tokenExpire"`
	// url prefix route depth
	RouteDepth int `yaml:"routeDepth"`
	// key : url prefix route value : route host list
	Route map[string]*RouteInfo `yaml:"routes"`
	// whitelist
	WhiteList []string `yaml:"whiteList"`
}

type RouteInfo struct {
	StripPrefix bool     `yaml:"stripPrefix"`
	Hosts       []string `yaml:"hosts"`
}

type Yaml struct {
	App appConfig `yaml:"gateway"`
}

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
}
