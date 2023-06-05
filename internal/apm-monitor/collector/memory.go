package collector

import (
	"APM-Monitor/pkg/tools"
	"github.com/shirou/gopsutil/v3/mem"
	"math"
)

func (c *CollectData) CollectMemoryData() {
	var metrics Metrics

	v, err := mem.VirtualMemory()
	if err != nil {
		c.errors <- &err
		return
	}
	metrics = Metrics{{
		Keys: map[string]string{
			"host":  Hostname(),
			"topic": c.topic,
		},
		Vals: map[string]interface{}{
			"total":       math.Round(float64(v.Total)/GBytes*100) / 100,
			"free":        math.Round(float64(v.Free)/GBytes*100) / 100,
			"used":        math.Round(float64(v.Used)/GBytes*100) / 100,
			"cached":      math.Round(float64(v.Cached)/GBytes*100) / 100,
			"usedPercent": math.Round(v.UsedPercent),
		},
		Timestamp: tools.NewTimeStamp(),
	}}
	c.metric <- metrics
}
