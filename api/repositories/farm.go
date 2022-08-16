package api

import (
	"github.com/borja.sueiro/blockchain-restful-api/blockchain"
	"github.com/borja.sueiro/blockchain-restful-api/models"
)

type FarmRepository struct {
	blockchain blockchain.Blockchain
}

func NewFarmRepository(blockchain blockchain.Blockchain) *FarmRepository {
	return &FarmRepository{blockchain}
}

func (repo *FarmRepository) GetFarms() []models.Farm {
	return repo.blockchain.GetFarms()
}

func (repo *FarmRepository) GetFarmById(id string) (models.Farm, error) {
	return repo.blockchain.GetFarmById(id)
}

func (repo *FarmRepository) AddNewFarm(newFarm models.Farm) error {
	return repo.blockchain.AddNewFarm(newFarm)
}

func (repo *FarmRepository) UpdateFarm(newFarmDefinition models.Farm) error {
	return repo.blockchain.UpdateFarm(newFarmDefinition)
}
