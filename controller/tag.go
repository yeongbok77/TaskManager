package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yeongbok77/TaskManager/logic"
	"go.uber.org/zap"
	"strconv"
)

// ApplyTagHandler 为 issue 分配 tag
func ApplyTagHandler(c *gin.Context) {
	var (
		issueId int64
		tagId   int64

		err error
	)
	// 获取 issueId
	if issueId, err = strconv.ParseInt(c.Query("issueId"), 0, 64); err != nil {
		zap.L().Error("ApplyTagHandler-->    strconv.ParseInt Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 获取 tagId
	if tagId, err = strconv.ParseInt(c.Query("tagId"), 0, 64); err != nil {
		zap.L().Error("ApplyTagHandler-->    strconv.ParseInt Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 业务处理
	if err = logic.ApplyTag(issueId, tagId); err != nil {
		zap.L().Error("ApplyTagHandler-->    logic.ApplyTag Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 操作成功
	ResponseSuccess(c, nil)
	return

}

// RemoveTagHandler  对 issue 附带的 tag 进行解绑
func RemoveTagHandler(c *gin.Context) {
	var (
		issueId int64
		tagId   int64
		err     error
	)

	// 获取 issueId 和 tagId 参数
	if issueId, err = strconv.ParseInt(c.Query("issueId"), 0, 64); err != nil {
		zap.L().Error("RemoveTagHandler-->    strconv.ParseInt Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	if tagId, err = strconv.ParseInt(c.Query("tagId"), 0, 64); err != nil {
		zap.L().Error("RemoveTagHandler-->    strconv.ParseInt Err:strconv.ParseInt Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 业务处理
	if err = logic.RemoveTag(issueId, tagId); err != nil {
		zap.L().Error("RemoveTagHandler-->    logic.RemoveTag Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 操作成功
	ResponseSuccess(c, nil)
	return

}

// ActionTagHandler tag 的增删改
func ActionTagHandler(c *gin.Context) {
	var (
		actionType int
		content    string
		err        error
	)
	// 获取 actionType 参数
	if actionType, err = strconv.Atoi(c.Query("actionType")); err != nil {
		zap.L().Error("ActionTagHandler-->    strconv.Atoi Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// actionType:  1 创建tag
	switch actionType {
	case 1:
		// 获取tag内容参数
		content = xssHander(c.Query("content"))
		if err = logic.CreateTag(content); err != nil {
			zap.L().Error("ActionTagHandler-->    logic.CreateTag Err:", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}

	}

	// 操作成功
	ResponseSuccess(c, nil)
	return

}
