package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// Conf 全局
var Conf *Config

// Config 解析配置文件
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	SessionSecret string `yaml:"session_secret"`
	GinMode       string `yaml:"gin_mode"`
	LogLevel      string `yaml:"log_level"`
	MysqlDsn      string `yaml:"mysql_dsn"`
	Redis         struct {
		Addr string `yaml:"addr"`
		Pass string `yaml:"pass"`
		Db   string `yaml:"db"`
	} `yaml:"redis"`
	Es struct {
		URL  string `yaml:"url"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	} `yaml:"es"`
}

// InitConf 初始化配置文件
func InitConf() {
	file, err := ioutil.ReadFile("configs/conf.yaml")
	if err != nil {
		log.Fatalf("打开配置文件异常: %v", err)
	}

	var c Config
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		log.Fatalf("解析配置文件异常: %v", err)
	}
	Conf = &c
}
