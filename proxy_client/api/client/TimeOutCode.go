package client

import (
	"github.com/1141520944/proxy/client/pkg/models"
	validateG "github.com/1141520944/proxy/client/pkg/util/validate"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 返回接口状态的Code
func ServerInitTimeOutCode(c *gin.Context) {
	result := new(models.ResultData)
	data := new(models.ClientInitRequest)
	if err := c.ShouldBindJSON(data); err != nil {
		zap.L().Error("ServerHandler with invalid fail", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			result.ResponseError(c, models.CodeServerBusy)
			return
		} else {
			result.ResponseErrorWithMsg(c, models.CodeInvalidParam, validateG.RemoveTopStruct(errs.Translate(validateG.Trans)))
			return
		}
	}
}
