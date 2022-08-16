// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/borja.sueiro/blockchain-restful-api/api"
	"github.com/borja.sueiro/blockchain-restful-api/api/controllers"
	api2 "github.com/borja.sueiro/blockchain-restful-api/api/repositories"
	"github.com/borja.sueiro/blockchain-restful-api/blockchain"
)

// Injectors from wire.go:

func InitializeApp() *api.RestfulApplication {
	blockchain := InitializeBlockchain()
	farmRepository := api2.NewFarmRepository(blockchain)
	farmController := controllers.NewFarmController(farmRepository)
	transportRepository := api2.NewTransportRepository(blockchain)
	transportController := controllers.NewTransportController(transportRepository)
	traceRepository := api2.NewTraceRepository(blockchain)
	traceController := controllers.NewTraceController(traceRepository)
	restfulApplication := api.NewApp(farmController, transportController, traceController)
	return restfulApplication
}

// wire.go:

func InitializeBlockchain() blockchain.Blockchain {
	return blockchain.NewHyperledgerApp()
}