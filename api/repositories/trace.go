package api

import (
	"github.com/borja.sueiro/blockchain-restful-api/blockchain"
	"github.com/borja.sueiro/blockchain-restful-api/models"
)

type TraceRepository struct {
	blockchain blockchain.Blockchain
}

func NewTraceRepository(blockchain blockchain.Blockchain) *TraceRepository {
	return &TraceRepository{blockchain}
}

func (repo *TraceRepository) GetTraces() []models.Trace {
	return repo.blockchain.GetTraces()
}

func (repo *TraceRepository) GetTraceById(id string) (models.Trace, error) {
	return repo.blockchain.GetTraceById(id)
}

func (repo *TraceRepository) AddNewTrace(newTrace models.Trace) error {
	return repo.blockchain.AddNewTrace(newTrace)
}

func (repo *TraceRepository) UpdateTrace(newTraceDefinition models.Trace) error {
	return repo.blockchain.UpdateTrace(newTraceDefinition)
}
