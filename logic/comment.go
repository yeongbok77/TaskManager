package logic

import (
	"context"
	"github.com/yeongbok77/TaskManager/dao/es"
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

	return
}
