package handler

import (
	"fmt"
	"gorm.io/gorm/clause"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/util"
)

// 订单处理

type FoodOrderItem struct {
	FoodMenuItem
	Count int // 数量
}

type FoodOrderItemReq struct {
	FoodList     []*FoodOrderItem `json:"food_list"`
	OpenId       string           `json:"open_id"`       // 用户信息
	TargetPeriod string           `json:"target_period"` // 订餐的目标时间，取值只有9个，前端已校验， 明天中午、后天早上、今天晚上等
}

// OrderFood 提交点餐订单
func OrderFood(w http.ResponseWriter, r *http.Request) {
	req := &FoodOrderItemReq{}
	err := util.BindJson(r, req)
	if err != nil {
		_, _ = fmt.Fprint(w, "绑定错误", err)
		return
	}
	openId, err := util.GetOpenIdFromHeader(r)
	if err != nil {
		_, _ = fmt.Fprint(w, "没有登录", err)
		return
	}
	req.OpenId = openId
	q1 := db.DB.OrderDBModel
	// 冲突的话只更新订单内容
	if err = q1.Clauses(clause.OnConflict{DoUpdates: clause.AssignmentColumns([]string{q1.Info.ColumnName().String()})}).Create(&model.OrderDBModel{
		OpenID:       openId,
		UniqueCode:   util.GenerateUUID(),
		TargetPeriod: req.TargetPeriod,
		Info:         util.Convert2JSONString(req.FoodList),
	}); err != nil {
		_, _ = fmt.Fprint(w, "数据库错误", err)
		return
	}
	util.ReturnSuccessJSON(w, "成功保存/覆写订单")
}
