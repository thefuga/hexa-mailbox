// Package http has all implementations of an HTTP adapter. As all adapters, it
// will not know the domain. Rather, it will communicate through mailboxes (see channel package).
package http

import (
	"context"
	"encoding/json"
	"fmt"
)

type (
	// Sender interface, implemented by channel.Channel. This is the receiving mailbox
	// of the target service.
	Sender interface {
		Send(Payload)
	}

	// Receiver interface, implemented by channel.Channel. This is the sending mailbox
	// of the target service.
	Receiver interface {
		Receive(context.Context) (Payload, error)
	}

	// Payload is structurally equal to dto.DTO and model.Model, but has it's onw
	// method set and struct tags.
	Payload struct {
		Foo string `json:"foo"`
		Bar int    `json:"bar"`
	}

	// Handler would implement an http.HandlerFunc. It won't actually implemente it for
	// simplicity. It also holds the mailboxes.
	Handler struct {
		sender   Sender
		receiver Receiver
	}
)

func NewHandler(sender Sender, receiver Receiver) *Handler {
	return &Handler{
		sender:   sender,
		receiver: receiver,
	}
}

// Handle would be a typical http.HandlerFunc. For simplicity, it does not implement
// an actual http.HandlerFunc.
// Handle takes a serialized value out of the context, unmarshalls it and sends it
// to the handler output mailbox. Then it awaits for a response on it's own receiving
// mailbox.
func (h Handler) Handle(ctx context.Context) error {
	var p Payload

	if err := json.Unmarshal(ctx.Value("payload").([]byte), &p); err != nil {
		return err
	}

	h.sender.Send(p)

	v, _ := h.receiver.Receive(ctx)

	fmt.Printf("Received response %v from the service\n", v)

	return nil
}
