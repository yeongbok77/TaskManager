package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/yeongbok77/TaskManager/models"
	"reflect"
	"strconv"
)

// InsertIssue 向 es 中写入 issue
func InsertIssue(issueInfo *models.IssueInfo) (err error) {
	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 执行插入操作
	_, err = client.Index().
		Index("issueinfo"). // 索引名称
		BodyJson(issueInfo).
		Do(ctx) // 执行请求，需要传入一个上下文对象

	return
}

// Search 根据用户的搜索内容, 搜索issue
func Search(q string, ctx context.Context) (issueIds []string, err error) {
	var (
		searchRes *elastic.SearchResult
	)
	// 创建 bool 查询
	boolQuery := elastic.NewBoolQuery()
	// 创建 match 查询
	// 设置 boost 优先级
	// issueContent 最高
	matchQueryIssueContent := elastic.NewMatchQuery("issue_content", q).Boost(10)
	termQueryTags := elastic.NewMatchQuery("tags", q).Boost(2)
	matchQueryComments := elastic.NewMatchQuery("comments", q).Boost(1)

	boolQuery.Should(matchQueryIssueContent, termQueryTags, matchQueryComments)

	searchRes, err = client.Search().
		Index("issueinfo").
		Query(boolQuery).
		Do(ctx)
	for _, v := range searchRes.Each(reflect.TypeOf(models.IssueInfo{})) {
		u := v.(models.IssueInfo)
		issueIds = append(issueIds, strconv.Itoa(int(u.IssueId)))
	}
	return
}
