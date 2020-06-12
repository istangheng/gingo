package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
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

func init() {
	file, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}

	var c Config
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	Conf = &c
}
