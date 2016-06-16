package basictracer

import (
	"sync"

	"github.com/opentracing/opentracing-go"
)

// SpanMetadata holds the basic Span metadata.
type SpanMetadata struct {
	// A probabilistically unique identifier for a [multi-span] trace.
	TraceID uint64

	// A probabilistically unique identifier for a span.
	SpanID uint64

	// Whether the trace is sampled.
	Sampled bool

	// The span's associated baggage.
	baggageLock sync.Mutex
	Baggage     map[string]string // initialized on first use
}

// BaggageItem belongs to the opentracing.SpanMetadata interface
func (c *SpanMetadata) BaggageItem(key string) string {
	// TODO: if we want to support onBaggage, need a pointer to the bt.Span.
	//   s.onBaggage(canonicalKey, val)
	//   if s.trim() {
	//   	return s
	//   }

	c.baggageLock.Lock()
	defer c.baggageLock.Unlock()

	if c.Baggage == nil {
		return ""
	}
	return c.Baggage[key]
}

// SetBaggageItem belongs to the opentracing.SpanMetadata interface
func (c *SpanMetadata) SetBaggageItem(key, val string) opentracing.SpanMetadata {
	c.baggageLock.Lock()
	defer c.baggageLock.Unlock()
	if c.Baggage == nil {
		c.Baggage = make(map[string]string)
	}
	c.Baggage[key] = val
	return c
}
