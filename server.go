package main

import (
	"net/http"
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"fmt"
)

var (
	rpcDuration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "rpc_duration_milliseconds",
			Help: "RPC Latency distribution.",
		},
		[]string{"service"},
	)

	rpcDurationHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "rpc_duration_histogram_milliseconds",
			Help: "RPC Latency distribution.",
			Buckets: prometheus.LinearBuckets(0, 10, 10),
		},
	)
)

func init()  {
	prometheus.MustRegister(rpcDuration)
	prometheus.MustRegister(rpcDurationHistogram)
}

func main() {
	{
		var inline int = 3
		fmt.Println(inline)
	}

	http.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		return
	}))
	http.Handle("/apple", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		apple(w, r)
	}))
	http.Handle("/metrics", prometheus.Handler())

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func apple(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("apple"))
	return
}
