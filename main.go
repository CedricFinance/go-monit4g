package main

import (
	"fmt"
	"monit4g/huawei4g"
	"monit4g/metrics"
	"time"
)

func main() {
	cfg := LoadConfig()
	router := huawei4g.New(cfg.Router.Host)

	publisher := metrics.NewPublisher(cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	for {
		stats, err := router.GetTrafficStatistics()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", stats)

		publisher.SendMetrics(cfg.Router.Name, stats)

		time.Sleep(1 * time.Minute)
	}
}
