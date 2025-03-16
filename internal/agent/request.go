package agent

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/vgnlkn/metrix/internal/entity"
	"github.com/vgnlkn/metrix/internal/handler/api"
	"go.uber.org/zap"
)

type Client struct {
	host   string
	logger *zap.Logger
}

func NewClient(host string, logger *zap.Logger) Client {
	return Client{host, logger}
}

func (c Client) UpdateMetrics(gm entity.GaugeMetrics, cm entity.CounterMetrics) {
	log := func(msg, name, vType string) {
		c.logger.Info(msg,
			zap.String("name", name),
			zap.String("type", vType),
		)
	}

	update := func(name, value, vType string) {
		err := c.sendMetric(name, value, vType)
		if err != nil {
			log(fmt.Sprintf("error updating metric err is %v", err), name, value)
			return
		}

		log("successfully updated metric", name, value)
	}

	for name, value := range gm {
		update(name, fmt.Sprintf("%f", value), entity.TypeGauge)
	}

	for name, value := range cm {
		update(name, fmt.Sprintf("%d", value), entity.TypeCounter)
	}

}

func (c Client) sendMetric(name, value, vType string) error {
	m, err := api.NewMetricsFromString(name, vType, value)
	if err != nil {
		return fmt.Errorf("agentSendMetric: %v", err)
	}
	data := m.ToBytes()
	if data == nil {
		return fmt.Errorf("agentSendMetric: expected not nill response")
	}

	client := http.Client{}
	io := bytes.NewReader(*data)

	resp, err := client.Post(fmt.Sprintf("%s/update", c.host), "application/json", io)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http status %s", resp.Status)
	}

	return nil
}
