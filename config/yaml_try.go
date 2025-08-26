package config

import (
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server   ConfigServer   `yaml:"server"`
	Database ConfigDatabase `yaml:"database"`
	Redis    Redis          `yaml:"redis"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}
type ConfigServer struct {
	Ip   string `yaml:"ip"`
	Port int    `yaml:"port"`
}
type ConfigDatabase struct {
	Url  string `yaml:"url"`
	Port int    `yaml:"port"`
}
type ConfigDataBase struct {
	Code       int `yaml:"code"`
	Data       ConfigDataBaseINfo
	Error      string `yaml:"error"`
	UpdateTime int64  `yaml:"updatetime"`
}
type ConfigDataBaseINfo struct {
	Id         int    `yaml:"id"`
	Tag        string `yaml:"tag"`
	Name       string `yaml:"name"`
	Origin     string `yaml:"origin"`
	Content    string `yaml:"content"`
	Created_at string `yaml:"created_at"`
	Updated_at string `yaml:"updated_at"`
}

func yaml_try() {
	//读取yaml
	yamlData, err := os.ReadFile("config/conf.yaml")
	if err != nil {
		fmt.Println("读取yaml文件失败:", err)
		return
	}
	//创建结构体实例用于接收解析后的数据
	var config Config
	//解析yaml
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		fmt.Println("解析yaml文件失败:", err)
		return
	}
	//修改yaml
	//把server.port改为8888
	config.Server.Port = 8888
	//把database.url改为192.168.33.33
	config.Database.Url = "192.168.33.33"
	//保存yaml
	newData, err := yaml.Marshal(&config)
	if err != nil {
		fmt.Println("序列化yaml失败:", err)
		return
	}
	err = os.WriteFile("config/conf.yaml", newData, 0644)
	if err != nil {
		fmt.Println("保存yaml文件失败:", err)
		return
	}
	//成功
	fmt.Println("yaml文件操作成功")
}

var Conf *Config

func MockConfig() {
	//	 todo 假的后续 用你yaml解析的来 port 用string
	Conf = &Config{Server: ConfigServer{Port: 8888}, Redis: Redis{Host: "127.0.0.1", Port: "6379"}}
}
