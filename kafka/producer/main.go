package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func SendMessage(msg string) {
	// 消息TOPIC
	topic := "shopping"

	// kafka配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 封装消息
	message := &sarama.ProducerMessage{}
	message.Topic = topic
	message.Value = sarama.StringEncoder(msg)

	// 连接kafka
	var kafkaHosts = []string{"hdfs-host1:9092", "hdfs-host2:9092", "hdfs-host3:9092", "hdfs-host4:9092"}
	producer, err := sarama.NewSyncProducer(kafkaHosts, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer producer.Close()

	// 发送消息
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}

	// 发送结果
	fmt.Printf("[KAFKA-PRODUCER]pid:%v offset:%v\n", partition, offset)
}
