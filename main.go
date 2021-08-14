package main

import (
	_ "github.com/golang/mock/mockgen/model" // for mockgen usage
	"github.com/wgarunap/goconf"
	"github.com/wgarunap/xm-rest-api/config"
	"github.com/wgarunap/xm-rest-api/metrics"
	"github.com/wgarunap/xm-rest-api/server"
)

func main() {
	cfg := &config.Conf{}
	if err := goconf.Load(cfg); err != nil {
		panic(err)
	}

	go metrics.MetricsRouter(cfg)

	notify := server.Serve(cfg, server.NewRouter(cfg))
	<-notify

}
