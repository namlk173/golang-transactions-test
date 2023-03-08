package main

import (
	"banking/boostrap"
	"banking/handler"
	"banking/repository"
	route "banking/router"
	"banking/usecase"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	app := boostrap.NewApplication()
	accountRepository := repository.NewAccountRepository(app.Client.Database(app.Env.DBName), "account")
	accountUseCase := usecase.NewAccountUseCase(accountRepository, time.Duration(app.Env.ContextTimeout)*time.Second)
	accountHandler := handler.AccountHandler{
		AccountUseCase: accountUseCase,
		Ctx:            ctx,
	}
	transferRepository := repository.NewTransferRepository(app.Client.Database(app.Env.DBName), "transfer", "account")
	transferUseCase := usecase.NewTransferUseCase(transferRepository, time.Duration(app.Env.ContextTimeout)*time.Second)

	transferHandler := handler.TransferHandler{
		TransferUseCase: transferUseCase,
		Ctx:             ctx,
	}

	router := route.NewRouter(accountHandler, transferHandler)

	serve := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", app.Env.ServerAddress, app.Env.ServerPort),
		Handler:      router,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	fmt.Printf("Sever are running in port: %v\n", app.Env.ServerPort)
	log.Fatalln(serve.ListenAndServe())
}
