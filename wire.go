//go:build wireinject
// +build wireinject

package main

import (
	"github.com/borja.sueiro/blockchain-restful-api/api"
	"github.com/borja.sueiro/blockchain-restful-api/api/controllers"
	repo "github.com/borja.sueiro/blockchain-restful-api/api/repositories"
	"github.com/borja.sueiro/blockchain-restful-api/blockchain"
	"github.com/google/wire"
)

func InitializeBlockchain() blockchain.Blockchain {
	return blockchain.NewHyperledgerApp()
}

func InitializeApp() *api.RestfulApplication {
	panic(
		wire.Build(
			InitializeBlockchain,
			repo.NewTransportRepository,
			repo.NewTraceRepository,
			controllers.NewTransportController,
			controllers.NewTraceController,
			api.NewApp,
		),
	)
}
