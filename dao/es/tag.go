package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

// InsertTag 根据 issueId 向 es 中写入tag的内容
func InsertTag(issueId int64, tagContent string, ctx context.Context) (err error) {
	// 查询条件
	query := elastic.NewTermQuery("issue_id", issueId)

	// 脚本
	script := elastic.NewScript("ctx._source.tags.add(params.new_tag)").Params(
		map[string]interface{}{
			"new_tag": tagContent,
		},
	)

	// 执行
	_, err = client.UpdateByQuery("issueinfo").Query(query).Script(script).Refresh("true").Do(ctx)
	if err != nil {
		zap.L().Error("client.UpdateByQuery Err:", zap.Error(err))
	}
	return
}

// RemoveTag 移除 issue 的 tag
func RemoveTag(issueId int64, tagContent string) (err error) {
	// 查询条件
	boolQuery := elastic.NewBoolQuery().
		Filter(elastic.NewTermQuery("issue_id", issueId))

	// 删除脚本
	scriptStr := "ctx._source.tags.remove(ctx._source.tags.indexOf(params.tags))"
	script := elastic.NewScript(scriptStr).Params(
		map[string]interface{}{
			"tags": tagContent,
		})

	// 执行
	_, err = client.UpdateByQuery("issueinfo").
		Query(boolQuery).Script(script).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		zap.L().Error("client.UpdateByQuery Err:", zap.Error(err))
	}
	return
}
