package collector

import (
	"APM-Monitor/pkg/tools"
	"github.com/shirou/gopsutil/v3/net"
	"math"
)

func (c *CollectData) CollectNetworkData() {
	var metrics Metrics
	io, err := net.IOCounters(false)
	if err != nil {
		c.errors <- &err
		return
	}

	for _, stat := range io {
		//fmt.Println(stat)
		metrics = append(metrics, Metric{
			Keys: map[string]string{
				"host":  Hostname(),
				"topic": c.topic,
			},
			Vals: map[string]interface{}{
				"send":       math.Round((float64(stat.BytesSent)/MBytes)*100) / 100,
				"recv":       math.Round((float64(stat.BytesRecv)/MBytes)*100) / 100,
				"packetSend": tools.ScientificToNumber(float64(stat.PacketsSent)),
				"packetRecv": tools.ScientificToNumber(float64(stat.PacketsRecv)),
			},
			Timestamp: tools.NewTimeStamp(),
		})
	}
	c.metric <- metrics
}
