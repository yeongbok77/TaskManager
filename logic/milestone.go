package logic

import (
	"github.com/yeongbok77/TaskManager/dao/mysql"
	"go.uber.org/zap"
)

// AddMilestone 为 issue 添加一个 milestone
func AddMilestone(issueId int64, content string) (err error) {
	if err = mysql.CreateMilestone(issueId, content); err != nil {
		zap.L().Error("mysql.AddMilestone Err", zap.Error(err))
	}
	return
}
