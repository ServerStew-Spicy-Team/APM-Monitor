package scheduler

import (
	"APM-Monitor/internal/apm-monitor/collector"
	"APM-Monitor/internal/apm-monitor/reporter"
	"APM-Monitor/internal/pkg/known"
	"APM-Monitor/internal/pkg/log"
	"context"
	"sync"
	"time"
)

func Schedule(ctx context.Context, topic string, wg *sync.WaitGroup) {
	defer wg.Done()
	c := collector.NewCollector(topic)

	for {
		switch c.GetTopic() {
		case known.CPU:
			c.CollectCPUData()
		case known.MEMORY:
			c.CollectMemoryData()
		case known.DISK:
			c.CollectDiskData()
		case known.NETWORK:
			c.CollectNetworkData()
		}
		select {
		case err := <-c.CollectorReturnError():
			log.Errorw("producer error", "err:", (*err).Error())
			continue
		case <-ctx.Done():
			log.Infow("collector exiting", "topic", c.GetTopic())
			return
		default:
			reporter.Report(c)
		}
		time.Sleep(5 * time.Second)
	}
}
