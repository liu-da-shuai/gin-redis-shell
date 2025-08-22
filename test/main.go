package main

import (
	"fmt"
	"io"

	//"log"
	"encoding/json"
	"gin-redis-shell/models"
	"net/http"
	"time"
)

// func main() {
// 	res, err := http.Get("http://api.xygeng.cn/one")
// 	checkError(err)
// 	data, err := io.ReadAll(res.Body)
// 	checkError(err)
// 	fmt.Printf("Got: %q", string(data))
// }

//	func checkError(err error) {
//		if err != nil {
//			log.Fatalf("Get : %v", err)
//		}
//	}
func GetDailyQuoteFromApi() (*models.QuoteResponse, error) {
	apiURL := "https://api.xygeng.cn/one" //c=a表示所有类型
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("http请求失败:%v", err)

	}
	defer resp.Body.Close()
	//检查响应状态
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API返回非200状态码:%d", resp.StatusCode)

	}
	//读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应体失败，%v", err)
	}
	var quote models.QuoteResponse
	//解析json
	err = json.Unmarshal(body, &quote)
	if err != nil {
		return nil, fmt.Errorf("解析json失败,%v", err)
	}
	return &quote, nil
}
func main() {
	quote, err := GetDailyQuoteFromApi()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Got: %q", quote)
}
