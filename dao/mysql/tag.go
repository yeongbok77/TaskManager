package mysql

import (
	"github.com/yeongbok77/TaskManager/models"
	"go.uber.org/zap"
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

// GetTagContent 根据 tagId 获取 tag 的内容
func GetTagContent(tagId int64) (content string, err error) {
	tag := &models.Tag{Id: tagId}
	if err = db.Select("content").Find(&tag).Error; err != nil {
		zap.L().Error("db.Select(\"content\").Find(&tag) Err:", zap.Error(err))
		return
	}
	return tag.Content, nil
}
