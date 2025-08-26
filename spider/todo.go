package spider

import (
	"encoding/json"
	"fmt"
	"gin-redis-shell/config"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// todo 这里去异步爬取数据
type Spider struct {
	site []Site
}

type Site struct {
	Url     string   `yaml:"url"`
	format  string   `yaml:"format"`
	dateKey string   `yaml:"dateKey"`
	mapper  []Mapper `yaml:"mapper"`
}

type Mapper struct {
	Name    string `yaml:"name"`
	Key     string `yaml:"key"`
	Default string `yaml:"default"`
}

func NewSpider() *Spider {
	conf := config.Conf.Spider
	site := make([]Site, 0)
	for _, s := range conf.SiteList {
		site = append(site, Site{
			Url:    s.Url,
			format: s.Format,
			//mapper: s.Mapper,
		})
	}
	return &Spider{
		site: site,
	}
}

func (s *Spider) Start() {
	for _, v := range s.site {
		go func() {
			defer func() {
				if err := recover(); err != nil {

				}
			}()
			v.crawl()
		}()
	}

}

type Result struct {
	Tag     string
	Name    string
	Content string
}

func (s *Site) crawl() {
	resp, err := http.DefaultClient.Get(s.Url)
	if err != nil {
		return
	}
	data := map[string]interface{}{}
	b, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &data)

	content, ok := data[s.dateKey].(string)
	// {"name":"水急客舟疾，山花拂面香。","from":"李白《秋浦歌十七首》"}
	if ok {
		return
	}
	r := Result{}
	for _, field := range s.mapper {
		// 定义正则表达式
		re := regexp.MustCompile(fmt.Sprintf("%s:,", field.Name))
		// 查找匹配的内容

		jsonStr := re.FindString(content)
		//"name":"水急客舟疾，山花拂面香。"
		str := strings.Split(jsonStr, ":")
		setFieldValue(r, field.Name, str[1])
	}

	time.Sleep(5 * time.Second)
	s.crawl()
}

func setFieldValue(obj interface{}, fieldName string, value interface{}) error {
	// 获取反射值对象
	val := reflect.ValueOf(obj)

	// 必须传入指针，否则无法设置值
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("必须传入结构体指针")
	}

	// 解引用指针，获取结构体值
	val = val.Elem()

	// 检查是否为结构体
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("指针指向的不是结构体")
	}

	// 获取字段
	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("字段 %s 不存在", fieldName)
	}

	// 检查字段是否可设置（必须是导出字段）
	if !field.CanSet() {
		return fmt.Errorf("字段 %s 不可设置（可能未导出）", fieldName)
	}

	// 将输入值转换为字段对应的类型
	valValue := reflect.ValueOf(value)
	if valValue.Type() != field.Type() {
		return fmt.Errorf("类型不匹配，字段类型 %s，值类型 %s", field.Type(), valValue.Type())
	}

	// 设置字段值
	field.Set(valValue)
	return nil
}
