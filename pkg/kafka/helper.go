package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
)

type Mypartition struct {
	PartitionMap map[string]int32
}

func (p *Mypartition) Partition(message *sarama.ProducerMessage, numPartitions int32) (int32, error) {
	key := message.Key.(sarama.StringEncoder)
	partition, ok := p.PartitionMap[string(key)]
	if !ok {
		// 如果映射关系中没有指定的键，则使用默认的分区选择逻辑
		partition, _ = sarama.NewRandomPartitioner(message.Topic).Partition(message, numPartitions)
	}

	return partition, nil
}

// 该方法的作用在下文源码分析中有详细解释
func (p *Mypartition) RequiresConsistency() bool {
	return true
}
func NewMyPartitioner(topic string) sarama.Partitioner {
	m := viper.GetStringMap("kafka.map")
	newMap := make(map[string]int32)
	for key, value := range m {
		switch v := value.(type) {
		case int:
			newMap[key] = int32(v)
		case int32:
			newMap[key] = v
		case int64:
			newMap[key] = int32(v)
		// 添加其他可能的类型转换
		default:
			fmt.Printf("无法将值转换为int32: %v\n", value)
		}
	}
	return &Mypartition{
		PartitionMap: newMap,
	}
}
