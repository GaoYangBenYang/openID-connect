package middleware

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// 读取配置文件
func init() {
	yamlFile, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Println(err.Error())
	}
}

// 全局配置文件
var Config *globalConfig

type globalConfig struct {
	Application  applicationConfig  `yaml:"Application"`
	MySQL        mysqlConfig        `yaml:"MySQL"`
	Redis        redisConfig        `yaml:"Redis"`
	Kafka        interface{}        `yaml:"Kafka"`
	JsonWebToken jsonWebTokenConfig `yaml:"JsonWebToken"`
}

type applicationConfig struct {
	Name string `yaml:"Name"`
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
}

type mysqlConfig struct {
	Address  string `yaml:"Address"`
	UserName string `yaml:"UserName"`
	Password string `yaml:"Password"`
	DBName   string `yaml:"DBName"`
}

type redisConfig struct {
	Address  string `yaml:"Address"`
	Password string `yaml:"Password"`
	DBName   string `yaml:"DBName"`
}

type jsonWebTokenConfig struct {
	HeaderAlg          string `yaml:"header_alg"`
	HeaderTyp          string `yaml:"header_typ"`
	PayloadIss         string `yaml:"payload_iss"`
	PayloadSub         string `yaml:"payload_sub"`
	PayloadAud         string `yaml:"payload_aud"`
	PayloadNbf         int    `yaml:"payload_nbf"`
	PayloadExp         int    `yaml:"payload_exp"`
	PayloadIat         int    `yaml:"payload_iat"`
	PayloadJti         string `yaml:"payload_jti"`
	SignatureSecretKey string `yaml:"signature_secret_key"`
}
