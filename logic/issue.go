package logic

import (
	"github.com/yeongbok77/TaskManager/dao/mysql"
	"github.com/yeongbok77/TaskManager/dao/redis"
	"github.com/yeongbok77/TaskManager/models"
	"go.uber.org/zap"
)

// ListIssue 列出所有 issue
func ListIssue(page, size int) (issues []*models.Issue, err error) {
	// 获取所有的 issue
	if issues, err = mysql.GetAllIssue(page, size); err != nil {
		zap.L().Error("mysql.GetAllIssue err", zap.Error(err))
		return
	}

	var (
		milestoneIds []string
		tagIds       []string
	)

	// 获取 issue 的 milestone 和 tag 以及 comment
	for i := range issues {
		// 获取 issue 的评论
		if err = mysql.GetComment(issues[i]); err != nil {
			zap.L().Error("mysql.GetComment Err:", zap.Error(err))
			return
		}

		// 获取 issue 的 milestone
		if milestoneIds, err = redis.GetMilestoneIds(issues[i].Id); err != nil {
			zap.L().Error("redis.GetMilestoneIds Err:", zap.Error(err))
			return
		}
		// 如果这个 issue 有 milestone 则进行mysql查询, 否则跳过
		if len(milestoneIds) != 0 {
			if issues[i].Milestones, err = mysql.GetMilestones(milestoneIds); err != nil {
				zap.L().Error("mysql.GetMilestones Err:", zap.Error(err))
				return
			}
		}

		// 获取 issue 的 tag
		if tagIds, err = redis.GetTagIds(issues[i].Id); err != nil {
			zap.L().Error("redis.GetTagIds Err:", zap.Error(err))
			return
		}
		// 如果这个 issue 有 tag 则进行mysql查询, 否则跳过
		if len(tagIds) != 0 {
			if issues[i].Tags, err = mysql.GetTags(tagIds); err != nil {
				zap.L().Error("mysql.GetTags Err:", zap.Error(err))
				return
			}
		}

	}

	return
}

// ActionAddIssue 添加一个 issue
func ActionAddIssue(content string) (err error) {
	if err = mysql.CreateIssue(content); err != nil {
		zap.L().Error("mysql.AddIssue Err:", zap.Error(err))
	}
	return
}

// ActionDeleteIssue 删除一个 issue
func ActionDeleteIssue(issueId int64) (err error) {
	if err = mysql.DeleteIssue(issueId); err != nil {
		zap.L().Error("mysql.DeleteIssue Err:", zap.Error(err))
	}
	return
}

// ActionUpdateIssue 修改一个 issue
func ActionUpdateIssue(issueId int64, content string) (err error) {
	if err = mysql.UpdateIssue(issueId, content); err != nil {
		zap.L().Error("mysql.UpdateIssue Err:", zap.Error(err))
	}
	return
}

// ListIssueTagFilter 根据 tag 对 issue 进行过滤查询
func ListIssueTagFilter(page, size int, tagIdsSlice []string) (issues []*models.Issue, err error) {
	var (
		IssueIdIntersection []string // issueId 的交集
	)

	// 取 tagId 所属 issueId 的交集
	if IssueIdIntersection, err = redis.GetIssueIntersection(tagIdsSlice); err != nil {
		zap.L().Error("redis.GetIssueIntersection Err:", zap.Error(err))
		return
	}

	// 没有数据 直接返回 nil 即可
	if IssueIdIntersection == nil {
		return nil, nil
	}

	// 根据 issueId 查询 issue
	if issues, err = mysql.GetIssuesPage(page, size, IssueIdIntersection); err != nil {
		zap.L().Error("mysql.GetIssues Err:", zap.Error(err))
		return
	}

	var (
		milestoneIds []string
		tagIds       []string
	)

	// 获取 issue 的 milestone 和 tag 以及 comment
	for i := range issues {
		// 获取 issue 的评论
		if err = mysql.GetComment(issues[i]); err != nil {
			zap.L().Error("mysql.GetComment Err:", zap.Error(err))
			return
		}

		// 获取 issue 的 milestone
		if milestoneIds, err = redis.GetMilestoneIds(issues[i].Id); err != nil {
			zap.L().Error("redis.GetMilestoneIds Err:", zap.Error(err))
			return
		}

		// 如果这个 issue 有 milestone 则进行mysql查询, 否则跳过
		if len(milestoneIds) != 0 {
			if issues[i].Milestones, err = mysql.GetMilestones(milestoneIds); err != nil {
				zap.L().Error("mysql.GetMilestones Err:", zap.Error(err))
				return
			}
		}

		// 获取 issue 的 tag
		if tagIds, err = redis.GetTagIds(issues[i].Id); err != nil {
			zap.L().Error("redis.GetTagIds Err:", zap.Error(err))
			return
		}

		// 如果这个 issue 有 tag 则进行mysql查询, 否则跳过
		if len(tagIds) != 0 {
			if issues[i].Tags, err = mysql.GetTags(tagIds); err != nil {
				zap.L().Error("mysql.GetTags Err:", zap.Error(err))
				return
			}
		}

	}

	return
}

// ListBasisMilestone 根据 milestoneId 列出 issues
func ListBasisMilestone(milestoneId string) (issues []*models.Issue, err error) {
	var (
		issueIds []string
	)

	// 查 redis 取出对应的 issueId
	if issueIds, err = redis.GetIssueIds(milestoneId); err != nil {
		zap.L().Error("redis.GetIssueIds Err:", zap.Error(err))
		return
	}

	// key 不存在, 返回 nil
	if issueIds == nil {
		return nil, nil
	}

	// mysql 获取 issue
	if issues, err = mysql.GetIssues(issueIds); err != nil {
		zap.L().Error("mysql.GetIssues Err:", zap.Error(err))
		return
	}

	return
}
