package user

import (
	"strconv"
	"time"

	"github.com/1141520944/proxy/server/dao/mysql"
	"github.com/1141520944/proxy/server/pkg/models"
	"github.com/1141520944/proxy/server/pkg/util/jwt"
	"github.com/1141520944/proxy/server/pkg/util/snowflake"
	validateG "github.com/1141520944/proxy/server/pkg/util/validate"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func (uh *UserHandler) LoginHandler(c *gin.Context) {
	sr := new(models.UserLoginRequest)
	result := new(models.ResultData)
	if err := c.ShouldBindJSON(sr); err != nil {
		zap.L().Error("LoginHandler()  with invalid fail", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			result.ResponseError(c, models.CodeServerBusy)
			return
		} else {
			result.ResponseErrorWithMsg(c, models.CodeInvalidParam, validateG.RemoveTopStruct(errs.Translate(validateG.Trans)))
			return
		}
	}
	code, err := uh.mysql.User_CheckByUsername(sr.UserName)
	if err != nil {
		zap.L().Error("uh.mysql.User_CheckByUsername() with fail-err", zap.Error(err))
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
	if code != nil {
		if code == mysql.ErrorUsernameNoExist {
			result.ResponseError(c, models.CodeUserNotExist)
			return
		} else {
			t, err := uh.mysql.User_SelectByUsernameOne(sr.UserName)
			if err != nil {
				zap.L().Error("uh.mysql.User_SelectByUsernameOne() with fail", zap.Error(err))
				result.ResponseError(c, models.CodeServerBusy)
				return
			}
			if t.Password != sr.Password {
				result.ResponseError(c, models.CodeInvalidPassword)
				return
			}
			t.PassLoginTime = time.Now()
			if err := uh.mysql.User_UpdateOne(t); err != nil {
				zap.L().Error("uh.mysql.User_UpdateOne() with fail", zap.Error(err))
				result.ResponseError(c, models.CodeServerBusy)
				return
			}
			token, _ := jwt.GenToken(t.UserID, t.Username)
			result.ResponseSuccess(c, gin.H{
				"user_id":  strconv.FormatInt(t.UserID, 10),
				"username": t.Username,
				"name":     t.Name,
				"token":    token,
			}, 1)
		}
	}
}
func (uh *UserHandler) RegisterHandler(c *gin.Context) {
	sr := new(models.UserRegisterRequest)
	result := new(models.ResultData)
	if err := c.ShouldBindJSON(sr); err != nil {
		zap.L().Error("RegisterHandler()  with invalid fail", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			result.ResponseError(c, models.CodeServerBusy)
			return
		} else {
			result.ResponseErrorWithMsg(c, models.CodeInvalidParam, validateG.RemoveTopStruct(errs.Translate(validateG.Trans)))
			return
		}
	}
	code, err := uh.mysql.User_CheckByUsername(sr.UserName)
	if err != nil {
		zap.L().Error("uh.mysql.User_CheckByUsername() with fail-err", zap.Error(err))
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
	if code != nil {
		if code == mysql.ErrorUsernameExist {
			result.ResponseError(c, models.CodeUsernameExist)
			return
		} else {
			user := &models.Table_user{
				Username:      sr.UserName,
				Password:      sr.Password,
				UserID:        snowflake.GenID(),
				PassLoginTime: time.Now(),
			}
			if err2 := uh.mysql.User_InsertOne(user); err2 != nil {
				zap.L().Error(" uh.mysql.User_InsertOne() with fail-err", zap.Error(err))
				result.ResponseError(c, models.CodeServerBusy)
				return
			} else {
				result.ResponseSuccess(c, models.CodeSuccess, 1)
			}
		}
	}
}
