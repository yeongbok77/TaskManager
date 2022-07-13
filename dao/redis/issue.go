package redis

import (
	"go.uber.org/zap"
)

func GetIssueIntersection(tagIds []string) (IssueIntersection []string, err error) {
	// 根据 tagId 获取 key, 并将 key 赋值到切片中
	for i, tagId := range tagIds {
		keyTagSet := getRedisKey(KeyTagSet) + tagId

		// 检查 key 是否存在
		// n > 0  表示存在, 否则不存在
		// 当 key 不存在时, 把这个 key 从切片里移除
		n, _ := rdb.Exists(keyTagSet).Result()
		if n > 0 {
			tagIds[i] = keyTagSet
		} else {
			if len(tagIds) == 1 {
				// 长度等于1, 说明这个切片里全都是不存在的key,直接返回 nil 即可
				// 如果再继续往下走, 就会panic, 所以直接返回.
				return nil, nil
			}
			tagIds = append(tagIds[:i], tagIds[i+1:]...)
		}

	}

	// 取交集
	if IssueIntersection, err = rdb.SInter(tagIds...).Result(); err != nil {
		zap.L().Error("rdb.SInter Err:", zap.Error(err))
	}
	return
}

// GetIssueIds 根据 milestoneId 获取 issueId
func GetIssueIds(milestoneId string) (issueIds []string, err error) {
	// 获取 milestoneSet 的 key
	keyMilestoneSet := getRedisKey(KeyMilestoneSet) + milestoneId

	// n > 0 表示key存在, 否则不存在
	n, _ := rdb.Exists(keyMilestoneSet).Result()
	if n <= 0 {
		return nil, nil
	}

	// 获取 issueIds
	if issueIds, err = rdb.SMembers(keyMilestoneSet).Result(); err != nil {
		zap.L().Error("rdb.SMembers Err:", zap.Error(err))
	}
	return

}
