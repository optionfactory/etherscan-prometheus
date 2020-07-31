package main

import (
	"etherscan-prometheus/collectors"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"flag"
)

func main() {
	flag.Usage = func() {
		const usage = "Usage: etherscan_exporter [option] [arg]\n\n"
		fmt.Fprint(flag.CommandLine.Output(), usage);
		flag.PrintDefaults()
		os.Exit(2)
	}

	bindAddress := flag.String("bind", ":9368", "listen address")
	apiKey := flag.String("apikey", "", "Etherscan API key")

	flag.Parse()
	if len(flag.Args()) > 0 {
		flag.Usage()
		os.Exit(1)
	}
	if (len(*apiKey)) == 0 {
		flag.Usage()
		os.Exit(1)
	}


	registry := prometheus.NewPedanticRegistry()
	registry.MustRegister(
		collectors.NewGasPriceCollector(*apiKey),
		collectors.NewCurrentBlockCollector(*apiKey),
	)

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		ErrorLog:      log.New(os.Stderr, log.Prefix(), log.Flags()),
		ErrorHandling: promhttp.ContinueOnError,
	})

	http.Handle("/metrics", handler)
	http.ListenAndServe(*bindAddress, nil)
}
