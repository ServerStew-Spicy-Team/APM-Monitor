package collector

import (
	"os"
)

const (
	KBytes = float64(1024)
	MBytes = float64(KBytes * 1024)
	GBytes = float64(MBytes * 1024)
	Mbytes = float64(MBytes / 8)
)

func Hostname() string {
	host, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return host
}

var _ Collector = (*CollectData)(nil)

type CollectData struct {
	topic  string
	errors chan *error
	metric chan Metrics
}

type Collector interface {
	GetTopic() string
	CollectorReturnError() <-chan *error
	CollectorReturnMetric() <-chan Metrics
	CollectCPUData()
	CollectMemoryData()
	CollectDiskData()
	CollectNetworkData()
}

func (c *CollectData) GetTopic() string {
	return c.topic
}

func (c *CollectData) CollectorReturnError() <-chan *error {
	return c.errors
}

func (c *CollectData) CollectorReturnMetric() <-chan Metrics {
	return c.metric
}
func NewCollector(topic string) Collector {
	c := CollectData{
		topic:  topic,
		errors: make(chan *error, 1),
		metric: make(chan Metrics, 1),
	}
	return &c
}
