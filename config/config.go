package config

import (
	"fmt"
	redis1 "gin-redis-shell/redis"
	"os"

	yaml "gopkg.in/yaml.v3"
)
var Conf *Config
type Config struct{
	Server ConfigServer `yaml:"server"`
	Redis ConfigRedis `yaml:"redis"`
	Database ConfigDatabase `yaml:"database"`
}
type ConfigServer struct{
	Ip string `yaml:"ip"`
	Port string `yaml:"port"`
}
type ConfigDatabase struct{
	Url string `yaml:"url"`
	Port int `yaml:"port"`
}
type ConfigDataBase struct{
	Code int `yaml:"code"`
	Data ConfigDataBaseINfo
	Error string `yaml:"error"`
	UpdateTime int64 `yaml:"updatetime"`
}
type ConfigRedis struct{
	Addr string `yaml:"addr"`
	Password string `yaml:"password"`
	db int `yaml:"db"`
}
type ConfigDataBaseINfo struct{
	Id int `yaml:"id"`
	Tag string `yaml:"tag"`
	Name string `yaml:"name"`
	Origin string `yaml:"origin"`
	Content string `yaml:"content"`				
	Created_at string `yaml:"created_at"`
	Updated_at string `yaml:"updated_at"`
}
func ParseConfig() *Config {
	Conf = &Config{Server: ConfigServer{Port: ":8080",Ip: "127.0.0.1"}, Redis: ConfigRedis{Addr: "127.0.0.1:6379",Password: "",db:0},}
	return Conf
}
func MockConfig() *Config {
  //读取yaml
  yamlData,err := os.ReadFile("config/conf.yaml")
  if err != nil{
	fmt.Println("读取yaml文件失败:",err)
   	panic(err)
}
//创建结构体实例用于接收解析后的数据

  //解析yaml
  err = yaml.Unmarshal(yamlData,&Conf)
  if err != nil{
	fmt.Println("解析yaml文件失败:",err)
   	panic(err)	
  }
  ParseConfig()
  Conf.SetDefault()
  redis1.InitRedis(Conf.Redis.Addr,Conf.Redis.Password,Conf.Redis.db)
  return Conf
}
func (config *Config) SetDefault() {
	if config.Server.Port == "" {
		config.Server.Port = ":8080"
	}
	if config.Server.Ip == "" {
		config.Server.Ip = "127.0.0.1"
	}
	if config.Redis.Addr == "" {
		config.Redis.Addr = "localhost:6379"
	}
}