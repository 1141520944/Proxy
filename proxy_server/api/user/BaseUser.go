package user

import (
	"strconv"

	"github.com/1141520944/proxy/server/pkg/models"
	"github.com/1141520944/proxy/server/pkg/util/request"
	validateG "github.com/1141520944/proxy/server/pkg/util/validate"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func (uh *UserHandler) UpdateInformationHandler(c *gin.Context) {
	sr := new(models.UserInformationRequest)
	result := new(models.ResultData)
	if err := c.ShouldBindJSON(sr); err != nil {
		zap.L().Error("UpdateInformationHandler()  with invalid fail", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			result.ResponseError(c, models.CodeServerBusy)
			return
		} else {
			result.ResponseErrorWithMsg(c, models.CodeInvalidParam, validateG.RemoveTopStruct(errs.Translate(validateG.Trans)))
			return
		}
	}
	//依据token拿到uid
	userID, err := request.GetCyrrentUserID(c)
	if err != nil {
		zap.L().Error("GetCyrrentUserID() failed", zap.Error(err))
		result.ResponseError(c, models.CodeNeedLogin)
		return
	}
	user := &models.Table_user{
		Username: sr.UserName,
		Name:     sr.Name,
		Password: sr.Password,
		UserID:   userID,
		Phone:    sr.Phone,
	}
	if err2 := uh.mysql.User_UpdateOne(user); err2 != nil {
		zap.L().Error(" uh.mysql.User_UpdateOne() failed", zap.Error(err))
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
	result.ResponseSuccess(c, models.CodeSuccess, 1)
}
func (uh *UserHandler) SelectAllHandler(c *gin.Context) {
	// p := new(models.ParamPage)
	result := new(models.ResultData)
	page, page_size, err := request.GetPageInfo(c)
	if err != nil {
		zap.L().Error("SelectPageInformationSuper() failed", zap.Error(err))
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
	re, count, err := uh.mysql.User_SelectAll(int(page), int(page_size))
	if err != nil {
		zap.L().Error("uh.mysql.User_SelectAll() failed", zap.Error(err))
		result.ResponseError(c, models.CodeServerBusy)
		return
	}
	users := make([]*models.UserInformationResponse, 0, len(re))
	for _, v := range re {
		user := &models.UserInformationResponse{
			Name:          v.Name,
			Username:      v.Username,
			Password:      v.Password,
			Phone:         v.Phone,
			PassLoginTime: v.PassLoginTime,
			ServerNum:     v.ServerNum,
			UserID:        strconv.FormatInt(v.UserID, 10),
		}
		users = append(users, user)
	}
	result.ResponseSuccessLayui(c, users, int(count))
}
func (uh *UserHandler) DeleteHandler(c *gin.Context) {
	uidStr := c.Param("uid")
	result := new(models.ResultData)
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		result.ResponseError(c, models.CodeInvalidParam)
		return
	}
	if err2 := uh.mysql.User_DeleteByIDOne(uid); err2 != nil {
		zap.L().Error(" uh.mysql.User_DeleteByIDOne() failed", zap.Error(err))
		result.ResponseError(c, models.CodeInvalidParam)
		return
	}
	result.ResponseSuccess(c, models.CodeSuccess, 1)
}
func (uh *UserHandler) SelectByidHandler(c *gin.Context) {
	uidStr := c.Param("uid")
	result := new(models.ResultData)
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		result.ResponseError(c, models.CodeInvalidParam)
		return
	}
	re, err2 := uh.mysql.User_SelectByIDOne(uid)
	if err2 != nil {
		zap.L().Error(" uh.mysql.User_DeleteByIDOne() failed", zap.Error(err))
		result.ResponseError(c, models.CodeInvalidParam)
		return
	}
	user := &models.UserInformationResponse{
		Name:          re[0].Name,
		Username:      re[0].Username,
		Password:      re[0].Password,
		Phone:         re[0].Phone,
		PassLoginTime: re[0].PassLoginTime,
		ServerNum:     re[0].ServerNum,
		UserID:        strconv.FormatInt(re[0].UserID, 10),
	}
	list := []*models.UserInformationResponse{user}
	result.ResponseSuccessLayui(c, list, 1)
}
