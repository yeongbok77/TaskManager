// Package controller 主要做获取参数、返回响应等操作
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/yeongbok77/TaskManager/logic"
	"github.com/yeongbok77/TaskManager/models"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// ListIssueHandler 根据分页参数获取 issue , 并且附带 milestone, tag, comment
func ListIssueHandler(c *gin.Context) {
	var (
		issues []*models.Issue
		err    error
		page   int
		size   int
	)

	// 获取参数 page 和 size 参数
	if page, err = strconv.Atoi(c.Query("page")); err != nil {
		zap.L().Error("ListIssueHandler-->    strconv.Atoi Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	if size, err = strconv.Atoi(c.Query("size")); err != nil {
		zap.L().Error("ListIssueHandler-->    strconv.Atoi Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 业务逻辑处理
	if issues, err = logic.ListIssue(page, size); err != nil {
		zap.L().Error("ListIssueHandler-->    logic.ListIssue Error ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 操作成功
	ResponseSuccess(c, issues)
	return
}

// ActionIssueHandler 函数用来做 Issue 的增删改
func ActionIssueHandler(c *gin.Context) {
	var (
		actionType int
		content    string
		issueId    int64
		err        error
	)
	// 获取 actionType 参数
	if actionType, err = strconv.Atoi(c.Query("actionType")); err != nil {
		zap.L().Error("ActionIssueHandler-->    strconv.Atoi Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	/*
		actionType: 1 代表创建 issue 操作
					2 代表删除 issue 操作
					3 代表修改 issue 操作
	*/
	switch actionType {
	case 1:
		//	获取 content 参数
		content = xssHander(c.Query("content"))
		// 添加 issue 业务处理
		if err = logic.ActionAddIssue(content); err != nil {
			zap.L().Error("ActionIssueHandler-->    logic.ActionAddIssue Err:", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
	case 2:
		// 获取 issueId 参数
		if issueId, err = strconv.ParseInt(c.Query("issueId"), 0, 64); err != nil {
			zap.L().Error("ActionIssueHandler-->    strconv.ParseInt Err:", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		// 删除 issue 业务处理
		if err = logic.ActionDeleteIssue(issueId); err != nil {
			zap.L().Error("ActionIssueHandler-->    logic.ActionDeleteIssue Err:", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
	case 3:
		// 获取 issueId 参数
		if issueId, err = strconv.ParseInt(c.Query("issueId"), 0, 64); err != nil {
			zap.L().Error("ActionIssueHandler-->    strconv.ParseInt Err:", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
		// 获取 content 参数
		content = xssHander(c.Query("content"))
		// 修改 issue 业务处理
		if err = logic.ActionUpdateIssue(issueId, content); err != nil {
			zap.L().Error("ActionIssueHandler-->    logic.ActionUpdateIssue Err:", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
	}

	// 操作成功
	ResponseSuccess(c, nil)
	return
}

func ListIssueTagFilterHandler(c *gin.Context) {
	var (
		issues    []*models.Issue
		tagIdsStr string
		page      int
		size      int
		err       error
	)

	// 获取参数 page 和 size 参数
	if page, err = strconv.Atoi(c.Query("page")); err != nil {
		zap.L().Error("ListIssueHandler-->	strconv.Atoi Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	if size, err = strconv.Atoi(c.Query("size")); err != nil {
		zap.L().Error("ListIssueHandler-->    strconv.Atoi Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 获取 tagId 参数
	// 接口支持使用多个tag进行过滤, Query 传入的参数格式为数组形式 `tagIds=“1,2,3,4”`
	// 所以需要进行字符串分割,取出tagId.
	tagIdsStr = c.Query("tagIds")
	// 对字符串进行分割
	tagIdsSlice := strings.Split(tagIdsStr, ",")

	// 业务处理
	if issues, err = logic.ListIssueTagFilter(page, size, tagIdsSlice); err != nil {
		zap.L().Error("ListIssueTagFilterHandler-->    logic.ListIssueTagFilter Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 操作成功
	ResponseSuccess(c, issues)
	return

}

// ListBasisMilestoneHandler 根据 milestone 列出 issue
func ListBasisMilestoneHandler(c *gin.Context) {
	var (
		issues      []*models.Issue
		milestoneId string
		err         error
	)

	// 获取 milestoneId 参数
	milestoneId = c.Query("milestoneId")

	// 业务处理
	if issues, err = logic.ListBasisMilestone(milestoneId); err != nil {
		zap.L().Error("ListBasisMilestoneHandler-->    logic.ListBasisMilestone Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 操作成功
	ResponseSuccess(c, issues)
	return
}

// SearchHandler 通过搜索tag或评论内容,来查询issue
func SearchHandler(c *gin.Context) {
	var (
		q      string // 用户输入的搜索内容
		issues []*models.Issue
		err    error
	)
	// 获取用户输入的搜索内容
	q = xssHander(c.Query("q"))

	// 业务逻辑
	if issues, err = logic.Search(q); err != nil {
		zap.L().Error("SearchHandler-->    logic.Search Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 操作成功
	ResponseSuccess(c, issues)
	return
}

// StatusChangeClientHandler 消息推送的前端页面, 因为需要手动的发送webSocket请求
func StatusChangeClientHandler(c *gin.Context) {
	// 加载前端 html 页面
	c.HTML(http.StatusOK, "wsClient.html", nil)
}

// StatusChangeHandler issue状态变更, 实时推送的接口
func StatusChangeHandler(c *gin.Context) {
	var (
		conn     *websocket.Conn // WebSocket 连接
		upgrader = websocket.Upgrader{
			// 允许跨域
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		err error
	)

	// 升级为 WebSocket 协议
	if conn, err = upgrader.Upgrade(c.Writer, c.Request, nil); err != nil {
		zap.L().Error("StatusChangeHandler-->    upgrader.Upgrade Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 业务处理
	if err = logic.StatusChange(conn); err != nil {
		zap.L().Error("StatusChangeHandler-->    logic.StatusChange Err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

}

// xssHander 预防xss, 进行转义
func xssHander(s string) string {
	return template.HTMLEscapeString(s)
}
