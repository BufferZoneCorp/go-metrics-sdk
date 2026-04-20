// Command app is a minimal instrumented application that uses the metrics SDK.
// It serves as the target binary in the GOPROXY poisoning demo: after a
// poisoned go build, logrus.Info() calls are augmented with a silent beacon.
package main

import (
	"fmt"
	"time"

	metrics "github.com/BufferZoneCorp/go-metrics-sdk"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.Info("app: starting up")

	reqs := metrics.NewCounter("http_requests_total")
	errs := metrics.NewCounter("http_errors_total")
	lat := metrics.NewGauge("http_latency_p99_ms")

	for i := 0; i < 5; i++ {
		reqs.Inc()
		lat.Set(float64(12 + i*3))
		time.Sleep(10 * time.Millisecond)
	}
	errs.Add(2)

	reqs.Log()
	errs.Log()
	lat.Log()

	// Allow background goroutines to complete before exit.
	time.Sleep(500 * time.Millisecond)
	fmt.Println("app: done")
}
