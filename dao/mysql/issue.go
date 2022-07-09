package mysql

import "github.com/yeongbok77/TaskManager/models"

// GetAllIssue 按分页参数批量获取 issue
func GetAllIssue(page, size int) (issues []*models.Issue, err error) {
	// select id, content from issues order by create_time desc limit 0, 5
	err = db.Select("id, content").Order("create_time desc").Limit(size).Offset((page - 1) * size).Find(&issues).Error
	return
}

// GetMilestone 根据 issue_id 获取 milestone
func GetMilestone(issueId int64) (milestones []*models.Milestone, err error) {
	err = db.Preload("Issue").Find(&milestones, "issue_id = ?", issueId).Error
	return
}

// GetTag 根据 issueId 来获取 tag
func GetTag(issueId int64) (tags []*models.Tag, err error) {
	err = db.Preload("Issue").Find(&tags, "issue_id = ?", issueId).Error
	return
}

// GetComment 根据 issueId 来获取 comment
func GetComment(issueId int64) (comments []*models.Comment, err error) {
	err = db.Preload("Issue").Find(&comments, "issue_id = ?", issueId).Error
	return
}

// AddIssue 根据用户写入的内容，创建一个 issue
func AddIssue(content string) (err error) {
	return
}
