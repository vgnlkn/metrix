package client

import (
	"fmt"
	"net/http"

	"github.com/vgnlkn/metrix/internal/metrix"
)

type Client struct {
	host string
}

func NewClient(host string) Client {
	return Client{host}
}

func (c Client) UpdateMetrics(gm metrix.GaugeMetrics, cm metrix.CounterMetrics) {
	for name, value := range gm {
		c.sendMetric(name, fmt.Sprintf("%f", value), metrix.TypeGauge)
	}

	for name, value := range cm {
		c.sendMetric(name, fmt.Sprintf("%d", value), metrix.TypeCounter)
	}

}

func (c Client) sendMetric(name, value, vType string) error {
	client := http.Client{}
	resp, err := client.Post(fmt.Sprintf("%s/update/%s/%s/%s", c.host, vType, name, value), "text/plain", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected http status %s", resp.Status)
	}

	fmt.Println(resp.Status, resp.StatusCode)
	return nil
}
