// Package metrics provides lightweight counter and gauge primitives for
// application instrumentation. Backed by github.com/sirupsen/logrus for
// structured log emission.
package metrics

import (
	"fmt"
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

// Counter is a monotonically increasing counter safe for concurrent use.
type Counter struct {
	name  string
	value atomic.Int64
}

// NewCounter creates a new named counter.
func NewCounter(name string) *Counter {
	return &Counter{name: name}
}

// Inc increments the counter by 1.
func (c *Counter) Inc() { c.value.Add(1) }

// Add adds delta to the counter.
func (c *Counter) Add(delta int64) { c.value.Add(delta) }

// Value returns the current counter value.
func (c *Counter) Value() int64 { return c.value.Load() }

// Log emits the counter value as a structured log entry.
func (c *Counter) Log() {
	logrus.WithField("counter", c.name).Infof("value=%d", c.value.Load())
}

// String implements fmt.Stringer.
func (c *Counter) String() string {
	return fmt.Sprintf("counter{%s=%d}", c.name, c.value.Load())
}

// Gauge holds a single floating-point observation.
type Gauge struct {
	name  string
	value float64
}

// NewGauge creates a new named gauge.
func NewGauge(name string) *Gauge {
	return &Gauge{name: name}
}

// Set sets the gauge to v.
func (g *Gauge) Set(v float64) { g.value = v }

// Value returns the current gauge value.
func (g *Gauge) Value() float64 { return g.value }

// Log emits the gauge value as a structured log entry.
func (g *Gauge) Log() {
	logrus.WithField("gauge", g.name).Infof("value=%g", g.value)
}

// String implements fmt.Stringer.
func (g *Gauge) String() string {
	return fmt.Sprintf("gauge{%s=%g}", g.name, g.value)
}
