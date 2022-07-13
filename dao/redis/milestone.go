package redis

import (
	"go.uber.org/zap"
	"strconv"
)

// ApplyMilestone 为 issue 分配 milestone
func ApplyMilestone(issueId, milestoneId int64) (err error) {
	// 获取 issue 的 milestone集合key
	keyIssueMilestoneSet := getRedisKey(KeyIssueMilestoneSet) + strconv.Itoa(int(issueId))

	// 获取 milestone 的集合key
	keyMilestoneSet := getRedisKey(KeyMilestoneSet) + strconv.Itoa(int(issueId))

	pipe := rdb.TxPipeline()

	// 向 issue 的 milestone 集合中加入 milestoneId
	pipe.SAdd(keyIssueMilestoneSet, milestoneId)

	// 向 milestone 集合中加入 issueId
	pipe.SAdd(keyMilestoneSet, issueId)

	_, err = pipe.Exec()

	return
}

// GetMilestoneIds 获取 issue 的所有 milestone 的 id
func GetMilestoneIds(issueId int64) (milestoneIds []string, err error) {
	// 获取 key
	keyIssueMilestoneSet := getRedisKey(KeyIssueMilestoneSet) + strconv.Itoa(int(issueId))
	n, err := rdb.Exists(keyIssueMilestoneSet).Result()
	if err != nil {
		zap.L().Error("rdb.Exists Err:", zap.Error(err))
		return
	}

	// n > 0  说明 key 存在, 否则说明不存在
	if n > 0 {
		// 获取集合中的所有 milestoneId
		if milestoneIds, err = rdb.SMembers(keyIssueMilestoneSet).Result(); err != nil {
			zap.L().Error("rdb.SMembers Err:", zap.Error(err))
		}
	}

	return
}
