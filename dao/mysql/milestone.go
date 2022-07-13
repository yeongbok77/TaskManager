package mysql

import (
	"github.com/yeongbok77/TaskManager/models"
	"go.uber.org/zap"
)

func GetMilestones(milestoneIds []string) (milestones []*models.Milestone, err error) {
	if err = db.Find(&milestones, milestoneIds).Error; err != nil {
		zap.L().Error("db.Find Err:", zap.Error(err))
	}
	return
}

// CreateMilestone 创建 milestone
func CreateMilestone(content string) (err error) {
	milestone := &models.Milestone{Content: content}
	err = db.Select("content").Create(&milestone).Error
	return
}
