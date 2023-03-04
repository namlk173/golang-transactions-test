package boostrap

import (
	"banking/mongoImplement"
	"context"
	"fmt"
	"log"
	"time"
)

func NewMongoDatabase(env *Env) mongoImplement.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	databaseURL := fmt.Sprintf("mongodb://%v:%v", env.DBHost, env.DBPort)
	if env.DBUser != "" {
		databaseURL = fmt.Sprintf("mongodb://%v:%v@%v:%v", env.DBUser, env.DBPass, env.DBHost, env.DBPort)
	}

	client, err := mongoImplement.NewClient(databaseURL)
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.Connect(ctx); err != nil {
		log.Fatalln(err)
	}

	if err := client.Ping(ctx); err != nil {
		log.Fatalln(err)
	}

	return client
}

func CloseMongoDatabase(client mongoImplement.Client) {
	if client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		log.Fatalln(err)
	}

	log.Println("Disconnect to database")
}
