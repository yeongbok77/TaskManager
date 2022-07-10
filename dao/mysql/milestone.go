package mysql

import "github.com/yeongbok77/TaskManager/models"

// CreateMilestone 创建一个 milestone
func CreateMilestone(issueId int64, content string) (err error) {
	milestone := &models.Milestone{IssueId: issueId, Content: content}
	err = db.Select("issue_id", "content").Create(&milestone).Error
	return
}
