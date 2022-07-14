package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

// InsertTag 根据 issueId 向 es 中写入tag的内容
func InsertTag(issueId int64, tagContent string, ctx context.Context) (err error) {

	query := elastic.NewTermQuery("issue_id", issueId)
	script := elastic.NewScript("ctx._source.tags.add(params.new_tag)").Params(
		map[string]interface{}{
			"new_tag": tagContent,
		},
	)

	_, err = client.UpdateByQuery("issueinfo").Query(query).Script(script).Refresh("true").Do(ctx)
	if err != nil {
		zap.L().Error("client.UpdateByQuery Err:", zap.Error(err))
	}
	return
}
