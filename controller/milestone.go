package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yeongbok77/TaskManager/logic"
	"go.uber.org/zap"
	"strconv"
)

func ActionMilestoneHandler(c *gin.Context) {

}

// AddMilestoneHandler 为 issue 分配 milestone
func AddMilestoneHandler(c *gin.Context) {
	var (
		issueId int64
		content string
		err     error
	)
	// 获取参数
	content = c.Query("content")
	if issueId, err = strconv.ParseInt(c.Query("issueId"), 0, 64); err != nil {
		zap.L().Error("AddMilestoneHandler-->    strconv.ParseInt Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 业务处理
	if err = logic.AddMilestone(issueId, content); err != nil {
		zap.L().Error("AddMilestoneHandler-->    logic.AddMilestone Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 操作成功
	ResponseSuccess(c, nil)
	return

}
