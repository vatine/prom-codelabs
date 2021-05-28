package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type gaugeFunc func(prometheus.Gauge, time.Duration)

// Start a function generator with a specific gauge vector and a
// specific period. This will spawn a never-stopping goroutine.
func startGauge(f gaugeFunc, gv *prometheus.GaugeVec, dt time.Duration) {
	seconds := dt / time.Second
	lv := fmt.Sprintf("%ds", seconds)
	g := gv.With(prometheus.Labels{"period": lv})
	fmt.Printf("Starting gauge metrics generator with period %s\n", dt)
	go f(g, dt)
}

// Start a latency generator with a specific HistogramVec and a
// specific "requests per second". This will spawn a never-stopping
// goroutine.
func startHisto(hv *prometheus.HistogramVec, qps int) {
	fmt.Printf("Starting histogram generatop with %d qps\n", qps)
	lv := fmt.Sprintf("%d", qps)
	go latencyGen(hv.With(prometheus.Labels{"qps": lv}), qps)
}

// Generate a sine wave with the full period dt
func sineGen(g prometheus.Gauge, dt time.Duration) {
	tick := time.NewTicker(time.Second)
	start := time.Now()
	interval := float64(dt)
	g.Set(0.0)
	for range tick.C {
		t1 := math.Mod(float64(time.Since(start)), interval)
		t2 := t1 / interval

		g.Set(math.Sin(2 * math.Pi * t2))
	}
}

// Generate a square wave with the full period dt
func squareGen(g prometheus.Gauge, dt time.Duration) {
	tick := time.NewTicker(time.Second)
	start := time.Now()
	interval := float64(dt) / float64(time.Second)
	half := interval / 2.0
	g.Set(1.0)

	for range tick.C {
		now := float64(time.Since(start) / time.Second)
		t := math.Mod(now, interval)
		if t > half {
			g.Set(-1.0)
		} else {
			g.Set(1.0)
		}
	}
}

// Compute a fake latency metric that has an approximately
// real-looking distribution.
func fakeLatency(attempts int, p float64) float64 {
	rv := 0.0

	for attempts > 0 {
		rv += 0.0015
		if rand.Float64() <= p {
			attempts--
		}
	}

	return rv
}

func latencyGen(h prometheus.Observer, qps int) {
	rate := time.Second / time.Duration(qps)
	tick := time.NewTicker(rate)

	for range tick.C {
		h.Observe(fakeLatency(5, 0.3))
	}
}

func main() {
	port := flag.String("listen", ":8080", "Port to start http listener on.")
	flag.Parse()

	sines := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "sine",
	}, []string{"period"})

	squares := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "square",
	}, []string{"period"})

	latency := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "latency",
		Buckets: prometheus.LinearBuckets(0.005, 0.001, 150),
	}, []string{"qps"})

	prometheus.MustRegister(sines)
	prometheus.MustRegister(squares)
	prometheus.MustRegister(latency)

	startGauge(sineGen, sines, 127*time.Second)
	startGauge(sineGen, sines, 293*time.Second)
	startGauge(sineGen, sines, 691*time.Second)

	startGauge(squareGen, squares, 131*time.Second)
	startGauge(squareGen, squares, 307*time.Second)
	startGauge(squareGen, squares, 701*time.Second)

	startHisto(latency, 10)
	startHisto(latency, 1000)

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Listening failed:\n\t%v\n", err)
		os.Exit(1)
	}
}
