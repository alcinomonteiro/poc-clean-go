package filialrepository

import (
	"log"

	"filial-go/app/core/dto"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func (repository repository) Create(filialInput *dto.CreateFilialInput) error {
	session, err := repository.collection.Database().Client().StartSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(repository.ctx)

	wc := writeconcern.New(writeconcern.WMajority())
	txnOptions := options.Transaction().SetWriteConcern(wc)

	err = mongo.WithSession(repository.ctx, session, func(ctx mongo.SessionContext) error {
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
		logrus.Info("Inserted: ", result.InsertedID)
		return nil
	})

	if err != nil {
		if err := session.AbortTransaction(repository.ctx); err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}

	return nil
}
