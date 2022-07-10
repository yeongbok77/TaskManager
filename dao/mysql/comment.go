package mysql

import "github.com/yeongbok77/TaskManager/models"

// CreateComment 创建一个评论
func CreateComment(issueId int64, content string) (err error) {
	comment := &models.Comment{IssueId: issueId, Content: content}
	err = db.Select("issue_id", "content").Create(&comment).Error
	return
}
