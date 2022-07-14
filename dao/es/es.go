package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
	"log"
	"os"
	"time"
)

var client *elastic.Client

// 索引mapping定义
const mapping = `
{
  "mappings": {
    "properties": {
      "issue_id": {
        "type": "integer"
      },
      "issue_content": {
        "type": "text"
      },
      "tags": {
        "type": "keyword"
      },
      "comments": {
        "type": "text"
      }
    }
  }
}`

func Init() (err error) {
	// 创建ES client用于后续操作ES
	client, err = elastic.NewClient(
		// 设置ES服务地址，支持多个地址
		elastic.SetURL("http://localhost:9200"),
		// 设置基于http base auth验证的账号和密码
		//elastic.SetBasicAuth("user", "secret")
		// 启用gzip压缩
		elastic.SetGzip(true),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置请求失败最大重试次数
		elastic.SetMaxRetries(5),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	if err != nil {
		// Handle error
		zap.L().Error("Elastic Search connection failed", zap.Error(err))
		panic(err)
	}

	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 首先检测下issueInfo索引是否存在
	exists, err := client.IndexExists("issueinfo").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// weibo索引不存在，则创建一个
		_, err := client.CreateIndex("issueinfo").BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
	}

	return
}
