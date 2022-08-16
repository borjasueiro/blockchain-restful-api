package blockchain

import (
	"github.com/borja.sueiro/blockchain-restful-api/models"
)

type MockBlochain struct {
	farms []models.Farm
}

func NewMockBlockchain() *MockBlochain {
	mock := &MockBlochain{}
	mock.farms = []models.Farm{
		{ID: "Feiraco", Location: "Muxía"},
		{ID: "O Rosal", Location: "Pereira"},
	}

	return mock
}

func (b *MockBlochain) GetFarms() []models.Farm {
	return b.farms
}

func (b *MockBlochain) GetFarmById(id string) (models.Farm, error) {
	for _, e := range b.farms {
		if e.ID == id {
			return e, nil
		}
	}
	return models.Farm{}, &AssetNotFoundError{id}
}

func (b *MockBlochain) AddNewFarm(newFarm models.Farm) error {
	for _, e := range b.farms {
		if e.ID == newFarm.ID {
			return &AssetAlreadyExistsError{newFarm.ID}
		}
	}
	b.farms = append(b.farms, newFarm)
	return nil
}

func (b *MockBlochain) UpdateFarm(newFarmDefinition models.Farm) error {
	for i, e := range b.farms {
		if e.ID == newFarmDefinition.ID {
			b.farms[i] = newFarmDefinition
			return nil
		}
	}
	return &AssetNotFoundError{newFarmDefinition.ID}
}
