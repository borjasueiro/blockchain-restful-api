package api

import (
	"github.com/borja.sueiro/blockchain-restful-api/blockchain"
	"github.com/borja.sueiro/blockchain-restful-api/models"
)

type TransportRepository struct {
	blockchain blockchain.Blockchain
}

func NewTransportRepository(blockchain blockchain.Blockchain) *TransportRepository {
	return &TransportRepository{blockchain}
}

func (repo *TransportRepository) GetTransports() []models.Transport {
	return repo.blockchain.GetTransports()
}

func (repo *TransportRepository) GetTransportById(id string) (models.Transport, error) {
	return repo.blockchain.GetTransportById(id)
}

func (repo *TransportRepository) AddNewTransport(newTransport models.Transport) error {
	return repo.blockchain.AddNewTransport(newTransport)
}

func (repo *TransportRepository) UpdateTransport(newTransportDefinition models.Transport) error {
	return repo.blockchain.UpdateTransport(newTransportDefinition)
}
