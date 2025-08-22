package models

import "encoding/json"

type Quote struct {
	ID     int      `json:"id"`
	Text   string   `json:"text"`
	Name   string   `json:"name"`
	Tag    []string `json:"tag"`
	Source string   `json:"source,omitempty"`
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
		Error      string `json:"error"`
		UpdateTime string `json:"updatetime"`
	} `json:"data"`
	Error      json.RawMessage `json:"error"`
	UpdateTime json.RawMessage `json:"updateTime"`
}
