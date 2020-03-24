package main

import (
	"github.com/CedricFinance/go-monit4g/huawei4g"
	"github.com/CedricFinance/go-monit4g/metrics"
	"log"
	"time"
)

func main() {
	cfg := LoadConfig()
	router := huawei4g.New(cfg.Router.Host)

	publisher := metrics.NewPublisher(cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	for {
		stats, err := router.GetTrafficStatistics()
		if err != nil {
			log.Printf("Failed to get the traffic statistics: %q\n", err)
			time.Sleep(10 * time.Second)
			continue
		}

		log.Printf("%+v\n", stats)

		err = publisher.SendMetrics(cfg.Router.Name, stats)
		if err != nil {
			log.Printf("Failed to send the metrics: %q\n", err)
			time.Sleep(10 * time.Second)
			continue
		}

		time.Sleep(1 * time.Minute)
	}
}
