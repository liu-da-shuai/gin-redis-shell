package dto

type Resp struct {
	Code int         `json:"code"` // 业务状态码
	Msg  string      `json:"msg"`  // 信息
	Data interface{} `json:"data"` // 数据
}
type QuoteResp struct {
	Code int `json:"code"`
	Data struct {
		ID         int    `json:"id"`
		Tag        string `json:"tag"`
		Name       string `json:"name"`
		Origin     string `json:"origin"`
		Content    string `json:"content"`
		Created_at string `json:"created_at"`
		Updated_at string `json:"updated_at"`
	} `json:"data"`
	Error      string `json:"error"`
	UpdateTime int64  `json:"updatetime"`
}