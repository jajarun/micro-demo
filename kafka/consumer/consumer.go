package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
)

type consumerGroupHandler struct {
	name string
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for msg := range claim.Messages() {
		fmt.Printf("%s Message topic:%q partition:%d offset:%d  value:%s\n", h.name, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		// 手动确认消息
		sess.MarkMessage(msg, "")
	}
	return nil
}

// kafka consumer

func main() {
	consumer, err := sarama.NewConsumerGroup([]string{"0.0.0.0:9092"}, "g1", nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return
	}
	_ = consumer.Consume(context.Background(), []string{"web_log"}, consumerGroupHandler{name: "g1"})
	//partitionList, err := consumer.Partitions("web_log") // 根据topic取到所有的分区
	//if err != nil {
	//	fmt.Printf("fail to get list of partition:err%v\n", err)
	//	return
	//}
	//fmt.Println(partitionList)
	//for partition := range partitionList { // 遍历所有的分区
	//	// 针对每个分区创建一个对应的分区消费者
	//	pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetOldest)
	//	if err != nil {
	//		fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
	//		return
	//	}
	//	defer pc.AsyncClose()
	//	// 异步从每个分区消费信息
	//	go func(sarama.PartitionConsumer) {
	//		for msg := range pc.Messages() {
	//
	//			fmt.Printf("Partition:%d Offset:%d Key:%v Value:%s\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
	//			pc.Pause()
	//		}
	//	}(pc)
	//}
	select {}
}
