package models

type Quote struct {
	Code int    `json:"code"`
	Message string `json:"message"`
	Result struct {
		Name	string   `json:"name"`
		From   string   `json:"from"`
	}
}
type QuoteRequest struct {
	Tag  string `json:"tag,omitempty"`
	Name string `json:"name,omitempty"`
}
type QuoteResponse struct {
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
