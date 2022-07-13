package mysql

import (
	"github.com/yeongbok77/TaskManager/models"
)

// CreateTag 创建一个 tag

// GetTags 根据 tagId 获取 tag
func GetTags(tagIds []string) (tags []*models.Tag, err error) {
	err = db.Find(&tags, tagIds).Error
	return
}

// CreateTag 创建 tag
func CreateTag(content string) (err error) {
	tag := &models.Tag{Content: content}
	err = db.Select("content").Create(&tag).Error
	return
}
