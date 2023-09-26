package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func InitConfig() {
	yamlFile, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		log.Println(err.Error())
	}
}

// 全局配置文件
var Conf *Config

type Config struct {
	Application    Application `yaml:"Application"`
	MySQL          MySQL       `yaml:"MySQL"`
	Redis          Redis       `yaml:"Redis"`
	Kafka          interface{} `yaml:"Kafka"`
	Log            interface{} `yaml:"Log"`
	OpenIDProvider interface{} `yaml:"OpenIDProvider"`
}

type Application struct {
	Name string `yaml:"Name"`
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

type MySQL struct {
	Root     string `yaml:"root"`
	Password string `yaml:"password"`
	DBName   string `yaml:"DBName"`
}

type Redis struct {
	Address  string `yaml:"Address"`
	Password string `yaml:"Password"`
	DBName   int    `yaml:"DBName"`
}
