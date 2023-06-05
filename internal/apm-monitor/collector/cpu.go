package collector

import (
	"APM-Monitor/pkg/tools"
	"github.com/shirou/gopsutil/v3/cpu"
	"math"
)

func (c *CollectData) CollectCPUData() {
	var total float64
	var metrics Metrics

	timestat, err := cpu.Times(false)
	if err != nil {
		c.errors <- &err
		return
	}
	for i, core := range timestat {
		//core := timestat[0]
		total = core.User + core.Idle + core.System + core.Nice + core.Iowait + core.Irq + core.Softirq + core.Steal + core.Guest + core.GuestNice
		metrics = append(metrics, Metric{
			Keys: map[string]string{
				"host":  Hostname(),
				"topic": c.topic,
			},
			Vals: map[string]interface{}{
				"cpu":    i,
				"user":   math.Round(core.User/total*10000) / 100,
				"system": math.Round(core.System/total*10000) / 100,
				"idle":   math.Round(core.Idle/total*10000) / 100,
			},
			Timestamp: tools.NewTimeStamp(),
		})
	}
	core := timestat[0]
	total = core.User + core.Idle + core.System + core.Nice + core.Iowait + core.Irq + core.Softirq + core.Steal + core.Guest + core.GuestNice
	metrics = Metrics{{
		Keys: map[string]string{
			"host":  Hostname(),
			"topic": c.topic,
		},
		Vals: map[string]interface{}{
			"user":   math.Round(core.User/total*10000) / 100,
			"system": math.Round(core.System/total*10000) / 100,
			"idle":   math.Round(core.Idle/total*10000) / 100,
		},
		Timestamp: tools.NewTimeStamp(),
	}}
	c.metric <- metrics
}
