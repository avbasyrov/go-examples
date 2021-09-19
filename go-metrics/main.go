package main

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("listen-address", ":8080",
	"The address to listen on for HTTP requests.")

func main() {
	usersRegistered := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "users_registered",
		})
	prometheus.MustRegister(usersRegistered)

	usersOnline := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "users_online",
		})
	prometheus.MustRegister(usersOnline)

	go func() {
		for {
			usersRegistered.Inc() // or: Add(5)
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	go func() {
		for {
			for i := 0; i < 10000; i++ {
				usersOnline.Set(float64(i)) // or: Inc(), Dec(), Add(5), Dec(5)
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())

	log.Printf("Starting web server at %s\n", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Printf("http.ListenAndServer: %v\n", err)
	}
}
