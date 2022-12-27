package model

import "encoding/json"

type Check struct {
	UserName    string `json:"user_name"`
	Task        string `json:"task"`
	TaskResults Result `json:"results"`
}

type Result struct {
	Payload []json.RawMessage `json:"payload"`
	Results []any             `json:"results"`
}
