// Package channel knows all port formats and enable comunication between them,
// as long as they have the same structural type.
package channel

import (
	"context"

	"github.com/thefuga/silly-ports-adaters/dto"
	"github.com/thefuga/silly-ports-adaters/http"
)

type (
	// ChanT is the union of all equivalent structs used on all adapters. This is useful
	// to decouple the DTO from adapter logic, allowing for different method sets and
	// struct tags.
	// The compatibility of all structs will be ensured by the type system: modifying
	// the structure of any element of this union will result in a compile-time error.
	ChanT interface {
		http.Payload | dto.DTO
	}

	// Channel specifies the type of the channel as In, and the type will be converted to
	// as Out.
	Channel[In, Out ChanT] struct {
		c chan In
	}
)

func NewChannel[In, Out ChanT]() Channel[In, Out] {
	return Channel[In, Out]{c: make(chan In)}
}

// Send just sends in to the underlying channel.
func (c Channel[In, Out]) Send(in In) {
	c.c <- in
}

// Receive listens to the underlying until ctx is done. When the underlying
// channel receives, the received value is converted to Out and returned. The
// compatibility of In and Out is ensured by the type system.
func (c Channel[In, Out]) Receive(ctx context.Context) (Out, error) {
	select {
	case <-ctx.Done():
		return *new(Out), ctx.Err()
	case v := <-c.c:
		return Out(v), nil
	}
}
