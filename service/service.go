package service

import (
	"context"
	"fmt"

	"github.com/thefuga/silly-ports-adaters/dto"
)

type (
	// Receiver interface, implemented by channel.Channel. This is the receiving mailbox
	// of this service.
	Receiver interface {
		Receive(context.Context) (dto.DTO, error)
	}

	// Sender interface, implemented by channel.Channel. This is the receiving mailbox
	// of an adapter and is used by the service to send responses.
	Sender interface {
		Send(dto.DTO)
	}

	// Service is the implementation of an arbitrary service. All communication with
	// the adapters is done trhough the channels, using DTOs.
	Service struct {
		receiver Receiver
		sender   Sender
	}
)

func NewService(receiver Receiver, sender Sender) *Service {
	return &Service{
		receiver: receiver,
		sender:   sender,
	}
}

// Perform receives on the service's receiving mailbox, performs the service logic
// and sends responses to the adapter's receiving mailbox.
func (s Service) Perform(ctx context.Context) error {
	v, err := s.receiver.Receive(ctx)

	if err != nil {
		return err
	}

	fmt.Printf("Received request %v from adapter\n", v)

	s.sender.Send(v)

	return nil
}
