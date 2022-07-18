package logic

import (
	"context"
	"fmt"
	"github.com/yeongbok77/TaskManager/dao/es"
	"github.com/yeongbok77/TaskManager/dao/kafka"
	"github.com/yeongbok77/TaskManager/dao/mysql"
	"go.uber.org/zap"
)

// AddComment 给 issue 添加评论的业务处理
func AddComment(issueId int64, content string) (err error) {

	// 将评论内容写入 MySQL
	if err = mysql.CreateComment(issueId, content); err != nil {
		zap.L().Error("mysql.CreateComment Err:", zap.Error(err))
	}

	ctx := context.Background()
	// 向 es 写入 comment
	if err = es.InsertComment(issueId, content, ctx); err != nil {
		zap.L().Error("es.InsertComment Err:", zap.Error(err))
		return
	}

	// 构造消息
	message := fmt.Sprintf("issueId:%d 新增了一条评论", issueId)

	// 向 kafka 写入消息
	go kafka.SendStatusChangeMsg(message)

	return
}

// DeleteComment 删除评论的业务处理
func DeleteComment(issueId, commentId int64) (err error) {
	var commentContent string

	// 获取评论的内容, 并删除这个评论
	if commentContent, err = mysql.DeleteComment(commentId); err != nil {
		zap.L().Error("mysql.DeleteComment Err:", zap.Error(err))
		return
	}

	// 从 es 中删除评论
	if err = es.DeleteComment(issueId, commentContent); err != nil {
		zap.L().Error("es.DeleteComment Err:", zap.Error(err))
		return
	}

	// 构造消息
	message := fmt.Sprintf("issueId:%d 删除了一条评论", issueId)

	// 向 kafka 写入消息
	go kafka.SendStatusChangeMsg(message)

	return
}
