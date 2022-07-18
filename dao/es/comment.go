package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

// InsertComment 根据 issueId 将评论内容写入 es
func InsertComment(issueId int64, comment string, ctx context.Context) (err error) {
	query := elastic.NewTermQuery("issue_id", issueId)
	script := elastic.NewScript("ctx._source.comments.add(params.new_comment)").Params(
		map[string]interface{}{
			"new_comment": comment,
		},
	)
	_, err = client.UpdateByQuery("issueinfo").Query(query).Script(script).Refresh("true").Do(ctx)
	if err != nil {
		zap.L().Error("client.UpdateByQuery Err:", zap.Error(err))
	}
	return
}

// DeleteComment 根据 issueId 将评论内容进行删除
func DeleteComment(issueId int64, commentContent string) (err error) {
	// 查询条件
	boolQuery := elastic.NewBoolQuery().
		Filter(elastic.NewTermQuery("issue_id", issueId))

	// 删除脚本
	scriptStr := "ctx._source.comments.remove(ctx._source.comments.indexOf(params.comments))"
	script := elastic.NewScript(scriptStr).Params(
		map[string]interface{}{
			"comments": commentContent,
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
