package main

import (
	"banking/boostrap"
	"banking/model"
	"banking/repository"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	idFrom, _ := primitive.ObjectIDFromHex("6401bd3f2b7d54b91cd0e719")
	idTo, _ := primitive.ObjectIDFromHex("6401bd492b7d54b91cd0e71a")

	app := boostrap.NewInformation()
	defer app.CloseDBConnection()
	database := app.Client.Database(app.Env.DBName)
	transferRepository := repository.NewTransferRepository(database, "transfer", "account")
	transfer := model.TransferRequest{
		From:   idFrom,
		To:     idTo,
		Amount: 10000,
	}
	fmt.Println(transferRepository.Transfer(ctx, &transfer))
}
