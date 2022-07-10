package models

type Tag struct {
	Id      int64  `gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT;"  json:"-"`
	IssueId int64  `gorm:"index:idx_issueId;  type:bigint(20) NOT NULL"  json:"-"`
	Content string `gorm:"type:varchar(10) collate utf8mb4_general_ci NOT NULL"`
}
