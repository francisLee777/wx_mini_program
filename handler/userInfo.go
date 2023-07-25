package handler

import (
	"fmt"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/util"
)

// GetUserInfo 获取用户信息： 目前是用户手动输入的昵称和头像
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	openId, err := util.GetOpenIdFromHeader(r)
	if err != nil {
		_, _ = fmt.Fprint(w, "没有登录", err)
		return
	}
	q1 := db.DB.UserInfoDBModel
	userInfo, err := q1.Where(q1.OpenID.Eq(openId)).First()
	if err != nil {
		_, _ = fmt.Fprint(w, err)
		return
	}
	util.ReturnSuccessJSON(w, userInfo)
}

// SaveNickName 保存用户昵称
func SaveNickName(w http.ResponseWriter, r *http.Request) {
	openId, err := util.GetOpenIdFromHeader(r)
	if err != nil {
		_, _ = fmt.Fprint(w, "没有登录", err)
		return
	}
	nickname := r.FormValue("nickname")
	q1 := db.DB.UserInfoDBModel
	_, err = q1.Where(q1.OpenID.Eq(openId)).Update(q1.UserNickName, nickname)
	if err != nil {
		_, _ = fmt.Fprint(w, err)
		return
	}
	util.ReturnSuccessJSON(w, nil)
}
