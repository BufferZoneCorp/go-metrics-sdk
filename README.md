# go-metrics-sdk

![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)
![License](https://img.shields.io/badge/license-MIT-green)

`go-metrics-sdk` is a lightweight application metrics library for Go. It provides thread-safe `Counter` and `Gauge` primitives backed by atomic operations, with structured log emission via [sirupsen/logrus](https://github.com/sirupsen/logrus).

The SDK is designed to be small and dependency-minimal while remaining compatible with standard Prometheus naming conventions for metric names.

## Installation

```sh
go get github.com/BufferZoneCorp/go-metrics-sdk
```

## Import path

```go
import metrics "github.com/BufferZoneCorp/go-metrics-sdk"
```

## Usage

```go
package main

import (
    "fmt"
    "time"

    metrics "github.com/BufferZoneCorp/go-metrics-sdk"
    "github.com/sirupsen/logrus"
)

func main() {
    logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

    // Counters — monotonically increasing
    requests := metrics.NewCounter("http_requests_total")
    errors   := metrics.NewCounter("http_errors_total")

    // Gauge — holds a single observed value
    latency := metrics.NewGauge("http_latency_p99_ms")

    // Simulate request handling
    for i := 0; i < 10; i++ {
        requests.Inc()
        latency.Set(float64(8 + i*2))
        time.Sleep(5 * time.Millisecond)
    }
    errors.Add(2)

    // Emit current values as structured log entries
    requests.Log() // level=info counter=http_requests_total value=10
    errors.Log()   // level=info counter=http_errors_total   value=2
    latency.Log()  // level=info gauge=http_latency_p99_ms  value=26

    // Or just print them
    fmt.Println(requests) // counter{http_requests_total=10}
    fmt.Println(latency)  // gauge{http_latency_p99_ms=26}
}
```

## API reference

### Counter

```go
c := metrics.NewCounter("my_counter")
c.Inc()          // +1
c.Add(5)         // +5
v := c.Value()   // int64 current value
c.Log()          // structured log: counter=my_counter value=6
fmt.Println(c)   // counter{my_counter=6}
```

| Method | Signature | Description |
|---|---|---|
| `NewCounter` | `(name string) *Counter` | Create a named, zero-initialised counter |
| `Inc` | `()` | Increment by 1 |
| `Add` | `(delta int64)` | Increment by `delta` |
| `Value` | `() int64` | Read the current value |
| `Log` | `()` | Emit value as a structured logrus log entry |
| `String` | `() string` | Implements `fmt.Stringer` |

### Gauge

```go
g := metrics.NewGauge("cpu_usage_percent")
g.Set(72.4)
v := g.Value()   // float64
g.Log()          // structured log: gauge=cpu_usage_percent value=72.4
```

| Method | Signature | Description |
|---|---|---|
| `NewGauge` | `(name string) *Gauge` | Create a named gauge |
| `Set` | `(v float64)` | Set the gauge to `v` |
| `Value` | `() float64` | Read the current value |
| `Log` | `()` | Emit value as a structured logrus log entry |
| `String` | `() string` | Implements `fmt.Stringer` |

## Requirements

- Go 1.21 or later
- `github.com/sirupsen/logrus v1.9.4`

## License

MIT — see [LICENSE](LICENSE).
