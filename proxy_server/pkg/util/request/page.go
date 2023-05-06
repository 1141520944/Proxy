package request

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserIDkey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// GetCyrrentUserID 获取当前用户的ID
func GetCyrrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDkey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// GetPageInfo 处理分页请求的参数
func GetPageInfo(c *gin.Context) (page, page_size int64, err error) {
	//获得分页参数offer和limit
	pageStr := c.Query("page")
	page_sizeStr := c.Query("size")
	page, err = strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		page = 1
	}
	page_size, err = strconv.ParseInt(page_sizeStr, 10, 32)
	if err != nil {
		page_size = 10
	}
	return
}
