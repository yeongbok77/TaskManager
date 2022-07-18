package kafka

import (
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

// SendStatusChangeMsg 向 kafka 中发送状态变更的消息
func SendStatusChangeMsg(message string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "issue"
	msg.Value = sarama.StringEncoder(message)

	// 连接kafka, 新建一个异步的生产者
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		zap.L().Error("producer closed, err:", zap.Error(err))
		return
	}
	defer producer.Close()

	// 发送消息
	producer.Input() <- msg
	if err != nil {
		zap.L().Error("send msg failed, err:", zap.Error(err))
		return
	}

}
