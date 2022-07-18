package mysql

import "github.com/yeongbok77/TaskManager/models"

// GetAllIssue 按分页参数批量获取 issue
func GetAllIssue(page, size int) (issues []*models.Issue, err error) {
	// select id, content from issues order by create_time desc limit 0, 5
	err = db.Select("id, content").Order("create_time desc").Limit(size).Offset((page - 1) * size).Find(&issues).Error
	return
}

// CreateIssue 根据用户写入的内容，创建一个 issue
func CreateIssue(issue *models.Issue) (err error) {
	err = db.Create(&issue).Error
	return
}

// DeleteIssue 根据 issueId 删除 issue
func DeleteIssue(issueId int64) (err error) {
	err = db.Delete(&models.Issue{}, issueId).Error
	return
}

// UpdateIssue 根据 issueId 和用户输入的内容，更新 issue内容
func UpdateIssue(issueId int64, content string) (err error) {
	err = db.Model(&models.Issue{}).Where("id = ?", issueId).Update("content", content).Error
	return
}

func GetComment(issue *models.Issue) (err error) {
	err = db.Model(&models.Issue{}).Preload("Comments").Find(&issue).Error
	return
}

// GetIssuesPage 根据 issueId 取 issue, 并且分页
func GetIssuesPage(page, size int, IssueIntersection []string) (issues []*models.Issue, err error) {
	err = db.Find(&issues, IssueIntersection).Limit(size).Offset((page - 1) * size).Error
	return
}

// GetIssues 根据 issueId 获取 issues
func GetIssues(issueIds []string) (issues []*models.Issue, err error) {
	err = db.Find(&issues, issueIds).Error
	return
}
