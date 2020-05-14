package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	// 消息TOPIC
	topic := "shopping"

	// 连接kafka
	var kafkaHosts = []string{"hdfs-host1:9092", "hdfs-host2:9092", "hdfs-host3:9092", "hdfs-host4:9092"}
	consumer, err := sarama.NewConsumer(kafkaHosts, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	defer consumer.Close()

	// 消息分区
	partitions, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(partitions)

	// 各分区消费消息
	for partition := range partitions {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		// 获取消息
		wg.Add(1)
		go func(c sarama.PartitionConsumer) {
			for {
				select {
				case msg := <-c.Messages():
					fmt.Printf("[KAFKA-CONSUMER]msg offset: %d, partition: %d, timestamp: %s, value: %s\n", msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
				case err := <-c.Errors():
					fmt.Printf("err :%s\n", err.Error())
				}
			}
			wg.Done()
		}(pc)
	}
	wg.Wait()

}
