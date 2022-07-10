package logic

import (
	"github.com/yeongbok77/TaskManager/dao/mysql"
	"go.uber.org/zap"
)

// AddComment 给 issue 添加评论的业务处理
func AddComment(issueId int64, content string) (err error) {
	if err = mysql.CreateComment(issueId, content); err != nil {
		zap.L().Error("mysql.CreateComment Err:", zap.Error(err))
	}
	return
}
