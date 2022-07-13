package logic

import (
	"github.com/yeongbok77/TaskManager/dao/mysql"
	"github.com/yeongbok77/TaskManager/dao/redis"
	"go.uber.org/zap"
)

// ApplyMilestone 为 issue 分配一个 milestone 业务处理
func ApplyMilestone(issueId, milestoneId int64) (err error) {
	if err = redis.ApplyMilestone(issueId, milestoneId); err != nil {
		zap.L().Error("redis.ApplyMilestone Err:", zap.Error(err))
	}
	return
}

// CreateMilestone 创建 milestone
func CreateMilestone(content string) (err error) {
	if err = mysql.CreateMilestone(content); err != nil {
		zap.L().Error("mysql.CreateMilestone Err:", zap.Error(err))
	}
	return
}
