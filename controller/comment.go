package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yeongbok77/TaskManager/logic"
	"go.uber.org/zap"
	"strconv"
)

// AddCommentHandler 为 issue 添加评论
func AddCommentHandler(c *gin.Context) {
	var (
		issueId int64
		content string
		err     error
	)
	// 获取参数
	content = xssHander(c.Query("content"))
	if issueId, err = strconv.ParseInt(c.Query("issueId"), 0, 64); err != nil {
		zap.L().Error("AddCommentHandler-->    strconv.ParseInt Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 业务处理
	if err = logic.AddComment(issueId, content); err != nil {
		zap.L().Error("AddCommentHandler-->    logic.AddComment Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 操作成功
	ResponseSuccess(c, CodeServerBusy)
	return

}

// DeleteCommentHandler 删除评论的接口
func DeleteCommentHandler(c *gin.Context) {
	var (
		issueId   int64
		commentId int64
		err       error
	)
	// 获取参数
	if issueId, err = strconv.ParseInt(c.Query("issueId"), 0, 64); err != nil {
		zap.L().Error("DeleteCommentHandler-->    strconv.ParseInt Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	if commentId, err = strconv.ParseInt(c.Query("commentId"), 0, 64); err != nil {
		zap.L().Error("DeleteCommentHandler-->    strconv.ParseInt Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 业务处理
	if err = logic.DeleteComment(issueId, commentId); err != nil {
		zap.L().Error("DeleteCommentHandler-->    logic.DeleteComment", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 操作成功
	ResponseSuccess(c, nil)
	return
}
