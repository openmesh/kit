package sd

import (
	"io"
	"testing"

	"github.com/go-kit/log"
	"github.com/openmesh/kit/endpoint"
)

func BenchmarkEndpoints(b *testing.B) {
	var (
		ca      = make(closer)
		cb      = make(closer)
		cmap    = map[string]io.Closer{"a": ca, "b": cb}
		factory = func(instance string) (endpoint.Endpoint[interface{}, interface{}], io.Closer, error) {
			return endpoint.Nop, cmap[instance], nil
		}
		c = newEndpointCache(factory, log.NewNopLogger(), endpointerOptions{})
	)

	b.ReportAllocs()

	c.Update(Event{Instances: []string{"a", "b"}})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Endpoints()
		}
	})
}
