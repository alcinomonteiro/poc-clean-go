package filialmessage

import (
	"context"
	"fmt"
	"strings"

	"filial-go/app/core/domain"
	"filial-go/app/core/error/jsonerror"

	"github.com/segmentio/kafka-go"
)

func Consume(ctx context.Context, r *kafka.Reader, controller domain.FilialController) {
	for {
		msg, err := r.FetchMessage(ctx)
		if err != nil {
			fmt.Println("could not read message ", err.Error())
			continue
		}
		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))

		if err := controller.Create(msg.Value); err != nil {
			fmt.Println("could not process message ", err.Error())
			if !strings.Contains(err.Error(), jsonerror.JsonErrorCode) {
				continue
			}
		}

		if err := r.CommitMessages(ctx, msg); err != nil {
			fmt.Println("could not commit message ", err.Error())
			continue
		}
	}
}
