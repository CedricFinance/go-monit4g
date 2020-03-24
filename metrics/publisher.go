package metrics

import (
	"fmt"
	"github.com/CedricFinance/go-monit4g/huawei4g"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)

type MetricsPublisher struct {
	database string
	client   client.Client
}

func NewPublisher(host, port, database string) *MetricsPublisher {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: fmt.Sprintf("http://%s:%s", host, port),
	})

	if err != nil {
		log.Fatal(err)
	}

	return &MetricsPublisher{
		database: database,
		client:   c,
	}
}

func (p MetricsPublisher) SendMetrics(routerName string, stats *huawei4g.TrafficStatistics) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  p.database,
		Precision: "s",
	})
	if err != nil {
		return err
	}

	tags := map[string]string{"router": routerName}
	fields := map[string]interface{}{
		"current_connect_time":  stats.CurrentConnectTime,
		"current_download":      stats.CurrentDownload,
		"current_download_rate": stats.CurrentDownloadRate,
		"current_upload":        stats.CurrentUpload,
		"current_upload_rate":   stats.CurrentUploadRate,
		"total_upload":          stats.TotalUpload,
		"total_download":        stats.TotalDownload,
	}

	pt, err := client.NewPoint("network_usage", tags, fields, time.Now())
	if err != nil {
		return err
	}
	bp.AddPoint(pt)

	if err := p.client.Write(bp); err != nil {
		return err
	}

	return p.client.Close()
}
