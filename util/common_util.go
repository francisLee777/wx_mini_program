package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg,omitempty"`
	Data     interface{} `json:"data"`
}

func BindJson(r *http.Request, ptr interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(ptr)
}

// GetOpenIdFromHeader 从Header里面获取微信 openId
func GetOpenIdFromHeader(r *http.Request) (string, error) {
	openId := r.Header.Get("x-wx-openid")
	if openId != "" {
		return "", fmt.Errorf("缺少openId")
	}
	return openId, nil
}

// ReturnSuccessJSON 向resp中注入json返回值
func ReturnSuccessJSON(w http.ResponseWriter, res interface{}) {
	w.Header().Set("content-type", "application/json")
	marshal, _ := json.Marshal(JsonResult{Data: res})
	_, _ = w.Write(marshal)
}
