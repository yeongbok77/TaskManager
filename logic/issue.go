package logic

import (
	"github.com/yeongbok77/TaskManager/dao/mysql"
	"github.com/yeongbok77/TaskManager/models"
	"go.uber.org/zap"
)

// ListIssue 列出所有 issue
func ListIssue(page, size int) (issues []*models.Issue, err error) {
	// 获取所有的 issue
	if issues, err = mysql.GetAllIssue(page, size); err != nil {
		zap.L().Error("mysql.GetAllIssue err", zap.Error(err))
		return
	}

	// 获取 issue 的 milestone 和 tag 以及 comment
	for i := range issues {
		if err = mysql.GetAllInformation(issues[i]); err != nil {
			zap.L().Error("mysql.GetAllInformation Err:", zap.Error(err))
			return
		}
	}

	return
}

// ActionAddIssue 添加一个 issue
func ActionAddIssue(content string) (err error) {
	if err = mysql.CreateIssue(content); err != nil {
		zap.L().Error("mysql.AddIssue Err:", zap.Error(err))
	}
	return
}

// ActionDeleteIssue 删除一个 issue
func ActionDeleteIssue(issueId int64) (err error) {
	if err = mysql.DeleteIssue(issueId); err != nil {
		zap.L().Error("mysql.DeleteIssue Err:", zap.Error(err))
	}
	return
}

// ActionUpdateIssue 修改一个 issue
func ActionUpdateIssue(issueId int64, content string) (err error) {
	if err = mysql.UpdateIssue(issueId, content); err != nil {
		zap.L().Error("mysql.UpdateIssue Err:", zap.Error(err))
	}
	return
}
