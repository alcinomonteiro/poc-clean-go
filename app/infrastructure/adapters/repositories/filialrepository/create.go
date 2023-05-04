package filialrepository

import (
	"context"
	"fmt"

	"filial-go/app/core/dto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func (repository repository) Create(filialInput *dto.CreateFilialInput) error {

	wc := writeconcern.New(writeconcern.WMajority())
	txnOptions := options.Transaction().SetWriteConcern(wc)

	session, err := repository.collection.Database().Client().StartSession()
	if err != nil {
		panic(err)
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(context.Background(), session, func(ctx mongo.SessionContext) error {
		if err = session.StartTransaction(txnOptions); err != nil {
			return err
		}
		result, err := repository.collection.InsertOne(ctx, filialInput)
		if err != nil {
			return err
		}
		if err = session.CommitTransaction(ctx); err != nil {
			return err
		}
		fmt.Println(result.InsertedID)
		return nil
	})

	if err != nil {
		if err := session.AbortTransaction(context.TODO()); err != nil {
			panic(err)
		}
		panic(err)
	}

	return nil
}
