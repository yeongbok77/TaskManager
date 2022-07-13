package redis

// redis key 的定义
const (
	Prefix               = "taskmanager:"
	KeyIssueTagSet       = "issue:tag:"       // issue 拥有的 tag 的集合; 根据 issue_id 来区分
	KeyIssueMilestoneSet = "issue:milestone:" // issue 拥有的 milestone 的集合; 根据 issue_id 来区分
	KeyTagSet            = "tag:"             // tag 的 issue_id 集合; 根据 tag_id 来区分
	KeyMilestoneSet      = "milestone:"       // milestone 的 issue_id 集合; 根据 milestone_id 来区分
)

func getRedisKey(key string) string {
	return Prefix + key
}
