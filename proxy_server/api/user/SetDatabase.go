package user

import (
	"github.com/1141520944/proxy/server/dao/mysql"
	"github.com/1141520944/proxy/server/pkg/models"
	"go.uber.org/zap"
)

func (uh *UserHandler) AddServerNum(userid int64) bool {
	code, err := uh.mysql.User_CheckByUserID(userid)
	if err != nil {
		zap.L().Error("uh.mysql.User_CheckByUsername() with fail-err", zap.Error(err))
		return false
	}
	if code != nil {
		if code == mysql.ErrorUsernameExist {
			re, _ := uh.mysql.User_SelectByIDOne(userid)
			now := re[0].ServerNum + 1
			user := &models.Table_user{ServerNum: now, UserID: userid}
			uh.mysql.User_UpdateOne(user)
			return true
		} else {
			return false
		}
	}
	return false
}
func (uh *UserHandler) DoneServerNum(userid int64) bool {
	code, err := uh.mysql.User_CheckByUserID(userid)
	if err != nil {
		zap.L().Error("uh.mysql.User_CheckByUsername() with fail-err", zap.Error(err))
		return false
	}
	if code != nil {
		if code == mysql.ErrorUsernameExist {
			re, _ := uh.mysql.User_SelectByIDOne(userid)
			now := re[0].ServerNum - 1
			user := &models.Table_user{ServerNum: now, UserID: userid}
			uh.mysql.User_UpdateOne(user)
			return true
		} else {
			return false
		}
	}
	return false
}
