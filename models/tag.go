package models

type Tag struct {
	Id      int64  `gorm:"primaryKey; type:bigint(20) AUTO_INCREMENT;"  json:"-"`
	Content string `gorm:"type:varchar(20) collate utf8mb4_general_ci NOT NULL"`
}
