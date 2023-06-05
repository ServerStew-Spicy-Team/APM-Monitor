package collector

import (
	"encoding/json"
	"github.com/Shopify/sarama"
)

var _ sarama.Encoder = (*Metric)(nil)
var _ sarama.Encoder = (*Metrics)(nil)

type Metric struct {
	Keys      map[string]string      `json:"keys"`
	Vals      map[string]interface{} `json:"vals"`
	Timestamp string                 `json:"timestamp"`
}

type Metrics []Metric

func (m *Metric) Encode() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *Metric) Length() int {
	data, _ := m.Encode()
	return len(data)
}

func (m Metrics) Encode() ([]byte, error) {
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m Metrics) Length() int {
	data, _ := m.Encode()
	return len(data)
}
