package controller

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	result interface{} `json:"result"`
}
