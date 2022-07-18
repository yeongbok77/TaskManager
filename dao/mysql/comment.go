package mysql

import "github.com/yeongbok77/TaskManager/models"

// CreateComment 创建一个评论
func CreateComment(issueId int64, content string) (err error) {
	comment := &models.Comment{IssueId: issueId, Content: content}
	err = db.Select("issue_id", "content").Create(&comment).Error
	return
}

// DeleteComment 删除评论, 并且返回这个评论
func DeleteComment(commentId int64) (commentContent string, err error) {
	comment := &models.Comment{Id: commentId}

	// 启动事务
	tx := db.Begin()

	// 先获取评论内容
	if err = tx.First(&comment, commentId).Error; err != nil {
		tx.Rollback()
	}

	// 将查询到的评论内容赋值给返回值
	commentContent = comment.Content

	// 删除评论
	if err = tx.Delete(&comment, commentId).Error; err != nil {
		tx.Rollback()
	}

	tx.Commit()

	return
}
