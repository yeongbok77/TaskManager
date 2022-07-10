package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yeongbok77/TaskManager/logic"
	"go.uber.org/zap"
	"strconv"
)

// AddTagHandler 为 issue 分配 tag
func AddTagHandler(c *gin.Context) {
	var (
		issueId int64
		content string
		err     error
	)
	// 获取参数
	content = c.Query("content")
	if issueId, err = strconv.ParseInt(c.Query("issueId"), 0, 64); err != nil {
		zap.L().Error("AddTagHandler-->    strconv.ParseInt Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 业务处理
	if err = logic.AddTag(issueId, content); err != nil {
		zap.L().Error("AddTagHandler-->    logic.AddTag Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 操作成功
	ResponseSuccess(c, nil)
	return

}
