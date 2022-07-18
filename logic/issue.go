package logic

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/gorilla/websocket"
	"github.com/yeongbok77/TaskManager/dao/es"
	"github.com/yeongbok77/TaskManager/dao/mysql"
	"github.com/yeongbok77/TaskManager/dao/redis"
	"github.com/yeongbok77/TaskManager/models"
	"go.uber.org/zap"
	"sync"
	"time"
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
	issue := &models.Issue{Content: content, CreateTime: time.Now()}
	if err = mysql.CreateIssue(issue); err != nil {
		zap.L().Error("mysql.AddIssue Err:", zap.Error(err))
		return
	}

	// 构造 es 中的 issueInfo 结构体
	issueInfo := &models.IssueInfo{IssueId: issue.Id, IssueContent: content, Tags: []string{" "}, Comments: []string{" "}}
	// elastic 执行插入操作
	if err = es.InsertIssue(issueInfo); err != nil {
		zap.L().Error("es.InsertIssue Err:", zap.Error(err))
		return
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

// Search 根据用户输入的内容, 搜索issue
func Search(q string) (issues []*models.Issue, err error) {
	var (
		issueIds []string
	)

	// 做查询
	ctx := context.Background()
	if issueIds, err = es.Search(q, ctx); err != nil {
		zap.L().Error("es.Search Err:", zap.Error(err))
		return
	}

	if len(issueIds) != 0 {
		zap.L().Error("ES 查出来的不为空")
	} else {
		zap.L().Error("ES 查出来的是空的")
	}

	// 根据 issueId 查询 issue
	if issues, err = mysql.GetIssues(issueIds); err != nil {
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

// StatusChange issue状态变更时的业务处理
func StatusChange(conn *websocket.Conn) (err error) {
	var wg sync.WaitGroup

	// 首先, 创建一个消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		zap.L().Error("sarama.NewConsumer Err:", zap.Error(err))
		return
	}

	// 这是一个大的for循环, 让它一直重复进行消费消息和写入WebSocket的操作
	for {

		// 根据 topic 取到所有分区
		partitionList, err := consumer.Partitions("issue")
		if err != nil {
			zap.L().Error("consumer.Partitions Err:", zap.Error(err))
			continue
		}

		// 循环的消费消息,并通过 WebSocket 推送到前端
		for _, p := range partitionList {
			partitionConsumer, err := consumer.ConsumePartition("issue", p, sarama.OffsetNewest)
			if err != nil {
				zap.L().Error("consumer.ConsumePartition Err:", zap.Error(err))
				continue
			}
			wg.Add(1)
			go func() {
				for m := range partitionConsumer.Messages() {
					conn.WriteMessage(websocket.TextMessage, m.Value)
				}
				wg.Done()
			}()

		}
		wg.Wait()

	}

}
