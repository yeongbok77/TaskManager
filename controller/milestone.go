package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yeongbok77/TaskManager/logic"
	"go.uber.org/zap"
	"strconv"
)

// ActionMilestoneHandler milestone 的增删改操作
func ActionMilestoneHandler(c *gin.Context) {
	var (
		actionType int
		content    string
		err        error
	)
	// 获取 actionType 参数
	if actionType, err = strconv.Atoi(c.Query("actionType")); err != nil {
		zap.L().Error("ActionMilestoneHandler-->    strconv.Atoi Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// actionType:  1 创建tag
	switch actionType {
	case 1:
		content = c.Query("content")
		if err = logic.CreateMilestone(content); err != nil {
			zap.L().Error("ActionMilestoneHandler-->    logic.CreateMilestone Err:", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
	}

	// 操作成功
	ResponseSuccess(c, nil)
	return
}

// ApplyMilestoneHandler 为 issue 分配 milestone
func ApplyMilestoneHandler(c *gin.Context) {
	var (
		issueId     int64
		milestoneId int64
		err         error
	)

	// 获取 issueId
	if issueId, err = strconv.ParseInt(c.Query("issueId"), 0, 64); err != nil {
		zap.L().Error("ApplyMilestoneHandler-->    strconv.ParseInt Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 获取 milestoneId
	if milestoneId, err = strconv.ParseInt(c.Query("milestoneId"), 0, 64); err != nil {
		zap.L().Error("ApplyMilestoneHandler-->    strconv.ParseInt Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 业务处理
	if err = logic.ApplyMilestone(issueId, milestoneId); err != nil {
		zap.L().Error("ApplyMilestoneHandler-->    logic.ApplyMilestone Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 操作成功
	ResponseSuccess(c, nil)
	return

}
