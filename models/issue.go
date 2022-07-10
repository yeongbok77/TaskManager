package models

import "time"

type Issue struct {
	Id         int64        `gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT;" json:"-"`
	Content    string       `gorm:"type:varchar(50) collate utf8mb4_general_ci NOT NULL" json:"content,omitempty"`
	CreateTime time.Time    `json:"create_time,omitempty"`
	Milestones []*Milestone `gorm:"foreignKey:IssueId" json:"milestones,omitempty"`
	Tags       []*Tag       `gorm:"foreignKey:IssueId" json:"tags,omitempty"`
	Comments   []*Comment   `gorm:"foreignKey:IssueId" json:"comments,omitempty"`
}
