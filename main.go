package main

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/thefuga/silly-ports-adaters/channel"
	"github.com/thefuga/silly-ports-adaters/dto"
	"github.com/thefuga/silly-ports-adaters/http"
	"github.com/thefuga/silly-ports-adaters/service"
)

func main() {
	var wg sync.WaitGroup

	// the service will receive messages from serviceMailbox and the handler will
	// produce messages to it.
	serviceMailbox := channel.NewChannel[http.Payload, dto.DTO]()

	// the handler will receive messages from handlerMailbox and it will be used
	// by the service to send responses.
	handlerMailbox := channel.NewChannel[dto.DTO, http.Payload]()

	// resolves the service and the handler with the appropriate mailboxes.
	h := http.NewHandler(serviceMailbox, handlerMailbox)
	s := service.NewService(serviceMailbox, handlerMailbox)

	{
		// Runs the service once until the service performs one action. Tipically
		// this would run during the whole application's lifecycle.
		wg.Add(1)
		go func() {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			s.Perform(ctx)
		}()
	}

	{
		// Runs the handler once. Tipically on an actual HTTP server, this would run
		// be invoked by the server for each request.
		wg.Add(1)
		go func() {
			defer wg.Done()
			payload := http.Payload{Foo: "foo", Bar: 1}
			bytes, err := json.Marshal(payload)

			if err != nil {
				return
			}

			ctx := context.WithValue(context.Background(), "payload", bytes)
			h.Handle(ctx)
		}()
	}

	// awaits for the handler and service goroutines.
	wg.Wait()
}
