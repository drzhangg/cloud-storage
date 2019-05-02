package util

import (
	"encoding/json"
	"fmt"
	"log"
)

//RespMsg：http响应数据的通用结构
type RespMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewRespMsg(code int, msg string, data interface{}) *RespMsg {
	return &RespMsg{
		code, msg, data,
	}
}

func (resp *RespMsg) JsonBytes() []byte {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return r
}

func (resp *RespMsg) JsonString() string {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return string(r)
}

//GenSimpleRespStream：只包含code和message的响应体([]byte)
func GenSimpleRespStream(code int, msg string) []byte {
	return []byte(fmt.Sprintf(`{"code":%d,"msg":%s}`, code, msg))
}

//GenSimpleRespString：只包含code和message的响应体(string)
func GenSimpleRespString(code int, msg string) string {
	return fmt.Sprintf(`{"code":%d,"msg":%s}`, code, msg)
}
