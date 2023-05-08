package main

import (
	"context"
	"net/http"

	"filial-go/app/infrastructure/adapters/messaging/filialmessage"
	"filial-go/app/infrastructure/config"
)

func main() {
	// context
	ctx := context.Background()
	// mongo client
	conn := config.ConnectMongo(ctx)
	defer conn.Disconnect(ctx)
	// mongo filial collection
	filialCollection := config.GetFilialCollection(conn)
	// di filial
	filialController := config.ConfigFilialDI(ctx, filialCollection)
	// kafka client filial
	filialReader := config.ConnectKafka()
	// Health
	go http.ListenAndServe(":8086", config.Health(conn, filialReader))
	// start listener
	filialmessage.Consume(ctx, filialReader, filialController)
}
