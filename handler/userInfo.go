package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"wxcloudrun-golang/util"
)

// GetUserPhoneNumber 获取用户手机号
func GetUserPhoneNumber(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	req := &struct {
		Code string `json:"code"`
	}{}
	err := util.BindJson(r, req)
	if err != nil {
		fmt.Fprint(w, "绑定错误", err)
		return
	}
	res.Data = r.Header
	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误", err)
		return
	}
	//decoder := json.
	marshal, _ := json.Marshal(req)
	httpResp, err := http.Post("https://api.weixin.qq.com/wxa/business/getuserphonenumber", "application/json", strings.NewReader(string(marshal)))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return
	}
	res.Data = string(body)
	res.ErrorMsg = string(marshal)
	msg, err = json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "内部错误", err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
