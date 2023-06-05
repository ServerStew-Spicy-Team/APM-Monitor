package collector

import (
	"APM-Monitor/pkg/tools"
	"github.com/shirou/gopsutil/v3/disk"
	"math"
	"time"
)

func (c *CollectData) CollectDiskData() {
	var metrics Metrics

	io, err := disk.IOCounters()
	if err != nil {
		c.errors <- &err
		return
	}
	for s, stat := range io {
		//fmt.Println(stat)
		metrics = append(metrics, Metric{
			Keys: map[string]string{
				"host":  Hostname(),
				"topic": c.topic,
			},
			Vals: map[string]interface{}{
				"disk":      s,
				"readIOPS":  tools.ScientificToNumber(float64(stat.ReadCount)),
				"writeIOPS": tools.ScientificToNumber(float64(stat.WriteCount)),
				"read":      math.Round((float64(stat.ReadBytes)/MBytes)*100) / 100,
				"write":     math.Round((float64(stat.WriteBytes)/MBytes)*100) / 100,
				//"readInSpeed":   math.Round(float64(stat.ReadCount) / (float64(stat.ReadTime) / float64(time.Second))),   // 个/s
				//"writeOutSpeed": math.Round(float64(stat.WriteCount) / (float64(stat.WriteTime) / float64(time.Second))), // 个/s
				"readFlow":  tools.ScientificToNumber(math.Round(((float64(stat.ReadBytes)/KBytes)/(float64(stat.ReadTime)/float64(time.Second)))*100) / 100),   // KB/s
				"writeFlow": tools.ScientificToNumber(math.Round(((float64(stat.WriteBytes)/KBytes)/(float64(stat.WriteTime)/float64(time.Second)))*100) / 100), // KB/s
			},
			Timestamp: tools.NewTimeStamp(),
		})
	}
	c.metric <- metrics
}
