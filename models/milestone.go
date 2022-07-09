package models

type Milestone struct {
	Id      int64  `gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT;"`
	IssueId int64  `gorm:"index:idx_issueId  type:bigint(20) not null"`
	Issue   Issue  `gorm:"foreignKey:IssueId"`
	Content string `gorm:"type:varchar(10) collate utf8mb4_general_ci NOT NULL"`
}
