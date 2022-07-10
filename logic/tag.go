package logic

import (
	"github.com/yeongbok77/TaskManager/dao/mysql"
	"go.uber.org/zap"
)

func AddTag(issueId int64, content string) (err error) {
	if err = mysql.CreateTag(issueId, content); err != nil {
		zap.L().Error("mysql.CreateTag Err:", zap.Error(err))
	}
	return
}
