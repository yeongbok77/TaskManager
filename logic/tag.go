package logic

import (
	"context"
	"fmt"
	"github.com/yeongbok77/TaskManager/dao/es"
	"github.com/yeongbok77/TaskManager/dao/kafka"
	"github.com/yeongbok77/TaskManager/dao/mysql"
	"github.com/yeongbok77/TaskManager/dao/redis"
	"go.uber.org/zap"
)

// ApplyTag 为 issue 分配 tag 业务处理
func ApplyTag(issueId, tagId int64) (err error) {
	var (
		tagContent string
	)
	// 为 issue 和 tag 的集合写入 id
	if err = redis.ApplyTag(issueId, tagId); err != nil {
		zap.L().Error("redis.ApplyTag Err:", zap.Error(err))
		return
	}

	// 获取 tag 的内容
	if tagContent, err = mysql.GetTagContent(tagId); err != nil {
		zap.L().Error("mysql.GetTagContent Err:", zap.Error(err))
		return
	}

	ctx := context.Background()
	// 向 es 写入 tagContent
	if err = es.InsertTag(issueId, tagContent, ctx); err != nil {
		zap.L().Error("es.InsertTag Err:", zap.Error(err))
		return
	}

	// 构造消息
	message := fmt.Sprintf("issueId:%d 与 tag:%s 进行绑定", issueId, tagContent)

	// 向 kafka 中写入消息
	go kafka.SendStatusChangeMsg(message)

	return
}

// RemoveTag	issue 解绑 tag 的业务处理
func RemoveTag(issueId, tagId int64) (err error) {

	// 从各自的 redis 集合中删除 id
	err = redis.RemoveTag(issueId, tagId)
	if err != nil {
		zap.L().Error("redis.RemoveTag Err:", zap.Error(err))
		return
	}

	// 从 mysql 中获取 tag 的内容, 以便让 es 进行删除操作
	tagContent, err := mysql.GetTagContent(tagId)
	if err != nil {
		zap.L().Error("mysql.GetTagContent Err:", zap.Error(err))
		return
	}

	// 从 es 中删除 issueId 附带的 tag
	err = es.RemoveTag(issueId, tagContent)
	if err != nil {
		zap.L().Error("es.RemoveTag Err:", zap.Error(err))
		return
	}

	// 构造消息
	message := fmt.Sprintf("issueId:%d 与 tag:%s 已解绑", issueId, tagContent)

	// 向 kafka 中写入 issue解绑tag 的消息
	go kafka.SendStatusChangeMsg(message)

	return
}

// CreateTag 创建 tag 的业务逻辑
func CreateTag(content string) (err error) {
	if err = mysql.CreateTag(content); err != nil {
		zap.L().Error("mysql.CreateTag Err:", zap.Error(err))
	}
	return
}
