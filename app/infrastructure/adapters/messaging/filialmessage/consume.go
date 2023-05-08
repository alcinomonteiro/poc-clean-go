package filialmessage

import (
	"context"
	"strings"

	"filial-go/app/core/domain"
	"filial-go/app/core/error/jsonerror"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func Consume(ctx context.Context, r *kafka.Reader, controller domain.FilialController) {
	for {
		msg, err := r.FetchMessage(ctx)
		if err != nil {
			logrus.Info("could not read message ", err.Error())
			continue
		}
		// after receiving the message, log its value
		logrus.Info("received: ", string(msg.Value))

		if err := controller.Create(msg.Value); err != nil {
			logrus.Info("could not process message ", err.Error())
			// Possibilidade de Identificação do erro para tratamento
			if !strings.Contains(err.Error(), jsonerror.JsonErrorCode) {
				continue
			}
		}

		if err := r.CommitMessages(ctx, msg); err != nil {
			logrus.Info("could not commit message ", err.Error())
		}
	}
}
