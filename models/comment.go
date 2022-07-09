package models

type Comment struct {
	Id      int64  `gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT"`
	Content string `gorm:"type:mediumtext collate utf8mb4_general_ci NOT NULL"`
	IssueId int64  `gorm:"index:idx_issueId type:bigint(20) not null"`
	Issue   Issue  `gorm:"foreignKey:IssueId"`
}
