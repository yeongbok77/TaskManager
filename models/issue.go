package models

import "time"

type Issue struct {
	Id         int64        `gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT;" json:"id,omitempty"`
	Content    string       `gorm:"type:varchar(50) collate utf8mb4_general_ci NOT NULL" json:"content,omitempty"`
	CreateTime time.Time    `json:"create_time,omitempty"`
	Milestones []*Milestone `gorm:"-" json:"milestones,omitempty"`
	Tags       []*Tag       `gorm:"-" json:"tags,omitempty"`
	Comments   []*Comment   `gorm:"-" json:"comments,omitempty"`
}
