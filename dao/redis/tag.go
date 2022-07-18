package redis

import (
	"go.uber.org/zap"
	"strconv"
)

// ApplyTag 为 issue 分配 tag
func ApplyTag(issueId, tagId int64) (err error) {
	// 获取 issue 的 tag集合key
	keyIssueTagSet := getRedisKey(KeyIssueTagSet) + strconv.Itoa(int(issueId))

	// 获取 tag 的集合key
	keyTagSet := getRedisKey(KeyTagSet) + strconv.Itoa(int(tagId))

	pipe := rdb.TxPipeline()

	// 向 issue 的 tag 集合中加入 tagId
	pipe.SAdd(keyIssueTagSet, tagId)

	// 向 tag 集合中加入 issueId
	pipe.SAdd(keyTagSet, issueId)

	_, err = pipe.Exec()

	return
}

// GetTagIds 获取 issue 的所有 tag 的 id
func GetTagIds(issueId int64) (tagIds []string, err error) {
	// 获取 key
	keyIssueTagSet := getRedisKey(KeyIssueTagSet) + strconv.Itoa(int(issueId))
	// 检查key是否存在
	n, err := rdb.Exists(keyIssueTagSet).Result()
	if err != nil {
		zap.L().Error("rdb.Exists Err:", zap.Error(err))
		return
	}

	// n > 0 说明key存在, 否则不存在
	if n > 0 {
		// 获取集合中的所有 tagId
		if tagIds, err = rdb.SMembers(keyIssueTagSet).Result(); err != nil {
			zap.L().Error("rdb.SMembers Err:", zap.Error(err))
		}
	}

	return
}

//// DeleteTag 删除 tag 对应的集合
//func DeleteTag(tagId int64) (issueIds []string, err error) {
//
//}
