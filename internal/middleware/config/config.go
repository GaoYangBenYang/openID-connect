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
	Application  application  `yaml:"Application"`
	MySQL        mySQL        `yaml:"MySQL"`
	Redis        redis        `yaml:"Redis"`
	Kafka        interface{}  `yaml:"Kafka"`
	JsonWebToken jsonWebToken `yaml:"JsonWebToken"`
}

type application struct {
	Name string `yaml:"Name"`
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

type mySQL struct {
	Address  string `yaml:"Address"`
	UserName string `yaml:"UserName"`
	Password string `yaml:"Password"`
	DBName   string `yaml:"DBName"`
}

type redis struct {
	Address  string `yaml:"Address"`
	Password string `yaml:"Password"`
	DBName   string `yaml:"DBName"`
}

type jsonWebToken struct {
	HeaderAlg       string `yaml:"header_alg"`
	HeaderTyp       string `yaml:"header_typ"`
	PayloadIss      string `yaml:"payload_iss"`
	PayloadSub      string `yaml:"payload_sub"`
	PayloadAud      string `yaml:"payload_aud"`
	PayloadNbf      int    `yaml:"payload_nbf"`
	PayloadExp      int    `yaml:"payload_exp"`
	PayloadIat      int    `yaml:"payload_iat"`
	PayloadJti      string `yaml:"payload_jti"`
	SignatureSecretKey string `yaml:"signature_secret_key"`
}
