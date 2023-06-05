package kafka

import (
	"github.com/Shopify/sarama"
	"sync"
)

var (
	once sync.Once
	// 全局变量，方便其它包直接调用已初始化好的 S 实例.
	p ProducerStore
)

type ProducerStore struct {
	P sarama.SyncProducer
}

type KafkaOptions struct {
	ProducerReturnSuccess bool
	ProducerReturnErr     bool
	Brokers               []string
}

func NewProducer(opts *KafkaOptions) (sarama.SyncProducer, error) {
	// 设置Kafka配置
	config := sarama.NewConfig()
	config.Producer.Return.Successes = opts.ProducerReturnSuccess
	config.Producer.Return.Errors = opts.ProducerReturnErr
	config.Producer.Partitioner = NewMyPartitioner

	client, err := sarama.NewClient(opts.Brokers, config)
	if err != nil {
		return nil, err
	}

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		//fmt.Println("Failed to create producer:", err)
		return nil, err
	}
	//defer producer.Close()

	return producer, nil
}

func StoreProducer(producer sarama.SyncProducer) {
	once.Do(func() {
		p = ProducerStore{producer}
	})
}

func Pro() ProducerStore {
	return p
}
