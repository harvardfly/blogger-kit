package config

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/ghodss/yaml"
)

// AppConfig 应用的配置结构体
type AppConfig struct {
	*ServerConfig `json:"server" yaml:"server"`
	*MySQLConfig  `json:"mysql" yaml:"mysql"`
	*RedisConfig  `json:"redis" yaml:"redis"`
	*LogConfig    `json:"log" yaml:"log"`
	*EtcdConfig   `json:"etcd" yaml:"etcd"`
}

// ServerConfig web server配置
type ServerConfig struct {
	Port int `json:"port" yaml:"port"`
}

// MySQLConfig MySQL数据库配置
type MySQLConfig struct {
	Host     string `json:"host" yaml:"host"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Port     string `json:"port" yaml:"port"`
	DB       string `json:"db" yaml:"db"`
	Debug    bool   `json:"debug" yaml:"debug"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host     string `json:"host" yaml:"host"`
	Password string `json:"password" yaml:"password"`
	Port     int    `json:"port" yaml:"port"`
	DB       int    `json:"db" yaml:"db"`
}

type LogConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      string
	Stdout     bool
}

// EtcdConfig etcd配置
type EtcdConfig struct {
	EtcdAddr string        `json:"etcdAddr" yaml:"etcdAddr"`
	SerName  string        `json:"serName" yaml:"serName"`
	GrpcAddr string        `json:"grpcAddr" yaml:"grpcAddr"`
	Ttl      time.Duration `json:"ttl" yaml:"ttl"`
}

// 定义了全局的配置文件实例
var Conf = new(AppConfig)

// Init 初始化
func Init(file string) error {
	yamlData, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := yaml.Unmarshal(yamlData, Conf); err != nil {
		return err
	}
	return nil
}
