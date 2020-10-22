package controller

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	result interface{} `json:"result"`
}
type ChatMessage struct {
	Time    string `json:"time"`
	Content string `json:"content"`
	MesType string `json:"mes_type"`
	Name    string `json:"name"`
	Avator  string `json:"avator"`
}
