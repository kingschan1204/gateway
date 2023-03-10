package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type AppConfig struct {
	// bind port
	Port string `yaml:"port"`
	// get validate code url
	Code string `yaml:"codeUrl"`
	// login url
	Login string `yaml:"loginUrl"`
	// url prefix route depth
	RouteDepth int `yaml:"routeDepth"`
	// key : url prefix route value : route host list
	Route map[string][]string `yaml:"routes"`
	// whitelist
	WhiteList []string `yaml:"whiteList"`
}

type AppYaml struct {
	App AppConfig `yaml:"gateway"`
}

var App AppConfig

func InitConfig() {
	//获取当前目录
	//fmt.Println(os.Getwd())

	filename := "./gateway.yaml"
	y := new(AppYaml)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read gateway.yaml file error %v\n", err)
	}
	err = yaml.Unmarshal(yamlFile, y)
	if err != nil {
		log.Fatalf("yaml 解码失败: %v\n", err)
	}
	App = y.App
}
