package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wxcloudrun-golang/util"
)

// 订单处理

type FoodOrderItem struct {
	FoodMenuItem
	Count int // 数量
}

type FoodOrderItemReq struct {
	FoodList []*FoodOrderItem `json:"food_list"`
	UserId   string           `json:"user_id"` // 用户信息？
}

// OrderFood 获取菜单列表接口
func OrderFood(w http.ResponseWriter, r *http.Request) {
	res := &JsonResult{}
	req := &FoodOrderItemReq{}
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
	w.Header().Set("content-type", "application/json")
	w.Write(msg)
}
