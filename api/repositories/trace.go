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

func (repo *TraceRepository) AddNewTrace(farm models.FarmRecollection) (string, error) {
	return repo.blockchain.AddNewTrace(farm)
}

func (repo *TraceRepository) AddFarmToTrace(id string, farm models.FarmRecollection) error {
	return repo.blockchain.AddFarmToTrace(id, farm)
}
func (repo *TraceRepository) AddTransvaseToTrace(id string, transvase models.Transvase) error {
	return repo.blockchain.AddTransvaseToTrace(id, transvase)
}
