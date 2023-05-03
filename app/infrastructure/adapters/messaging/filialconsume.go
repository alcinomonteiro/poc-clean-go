package messaging

import (
	"context"
	"fmt"

	"filial-go/app/core/domain"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context, r *kafka.Reader, controller domain.FilialController) {
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.FetchMessage(ctx)
		if err != nil {
			//panic("could not read message " + err.Error())
			fmt.Println("could not read message ", err.Error())
			continue
			// TODO: Precisamos considerar este panic.
			// utilizar outras estratégias
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))

		if err := controller.Create(msg.Value); err != nil {
			//panic("could not process message " + err.Error())
			fmt.Println("could not process message ", err.Error())
			continue
			// TODO: Precisamos considerar este panic.
			// utilizar outras estratégias
		}

		if err := r.CommitMessages(ctx, msg); err != nil {
			//panic("could not commit message " + err.Error())
			fmt.Println("could not commit message ", err.Error())
			continue
			// TODO: Precisamos considerar este panic.
			// utilizar outras estratégias
		}
	}
}
