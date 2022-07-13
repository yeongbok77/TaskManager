package logic

import (
	"github.com/yeongbok77/TaskManager/dao/mysql"
	"github.com/yeongbok77/TaskManager/dao/redis"
	"go.uber.org/zap"
)

// ApplyTag 为 issue 分配 tag 业务处理
func ApplyTag(issueId, tagId int64) (err error) {
	if err = redis.ApplyTag(issueId, tagId); err != nil {
		zap.L().Error("redis.ApplyTag Err:", zap.Error(err))
	}
	return
}

func CreateTag(content string) (err error) {
	if err = mysql.CreateTag(content); err != nil {
		zap.L().Error("mysql.CreateTag Err:", zap.Error(err))
	}
	return
}
