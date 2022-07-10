package mysql

import "github.com/yeongbok77/TaskManager/models"

// CreateTag 创建一个 tag
func CreateTag(issueId int64, content string) (err error) {
	tag := &models.Tag{IssueId: issueId, Content: content}
	err = db.Select("issue_id", "content").Create(&tag).Error
	return
}
