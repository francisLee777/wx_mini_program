package handler

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
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
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		_, _ = fmt.Fprint(w, "数据库错误", err)
		return
	}
	// 查不到时 userInfo 为空
	util.ReturnSuccessJSON(w, userInfo)
}

// SaveNickName 保存用户昵称
func SaveNickName(w http.ResponseWriter, r *http.Request) {
	openId, err := util.GetOpenIdFromHeader(r)
	if err != nil {
		_, _ = fmt.Fprint(w, "没有登录", err)
		return
	}
	nickname := r.URL.Query().Get("nickname")
	if nickname == "" {
		_, _ = fmt.Fprint(w, "入参nickName缺失")
		return
	}
	q1 := db.DB.UserInfoDBModel
	if err = q1.Clauses(clause.OnConflict{DoUpdates: clause.AssignmentColumns([]string{q1.UserNickName.ColumnName().String()})}).
		Create(&model.UserInfoDBModel{
			OpenID:       openId,
			UserNickName: nickname,
		}); err != nil {
		_, _ = fmt.Fprint(w, "数据库错误", err)
		return
	}
	util.ReturnSuccessJSON(w, nickname)
}

// SaveIconURL 保存用户头像链接
func SaveIconURL(w http.ResponseWriter, r *http.Request) {
	openId, err := util.GetOpenIdFromHeader(r)
	if err != nil {
		_, _ = fmt.Fprint(w, "没有登录", err)
		return
	}
	iconURL := r.URL.Query().Get("iconURL")
	if iconURL == "" {
		_, _ = fmt.Fprint(w, "入参iconURL缺失")
		return
	}
	q1 := db.DB.UserInfoDBModel
	if err = q1.Clauses(clause.OnConflict{DoUpdates: clause.AssignmentColumns([]string{q1.UserIconURL.ColumnName().String()})}).
		Create(&model.UserInfoDBModel{
			OpenID:      openId,
			UserIconURL: iconURL,
		}); err != nil {
		_, _ = fmt.Fprint(w, "数据库错误", err)
		return
	}
	util.ReturnSuccessJSON(w, iconURL)
}
