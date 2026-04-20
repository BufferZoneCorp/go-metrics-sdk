package metrics

import (
	"testing"
)

func TestCounter(t *testing.T) {
	c := NewCounter("requests")
	c.Inc()
	c.Inc()
	c.Add(3)
	if got := c.Value(); got != 5 {
		t.Fatalf("expected 5, got %d", got)
	}
	c.Log()
}

func TestCounterString(t *testing.T) {
	c := NewCounter("errors")
	c.Add(7)
	if s := c.String(); s != "counter{errors=7}" {
		t.Fatalf("unexpected string: %q", s)
	}
}

func TestGauge(t *testing.T) {
	g := NewGauge("latency_ms")
	g.Set(42.5)
	if got := g.Value(); got != 42.5 {
		t.Fatalf("expected 42.5, got %g", got)
	}
	g.Log()
}

func TestGaugeString(t *testing.T) {
	g := NewGauge("cpu")
	g.Set(0.75)
	if s := g.String(); s != "gauge{cpu=0.75}" {
		t.Fatalf("unexpected string: %q", s)
	}
}
