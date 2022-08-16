package blockchain

import "github.com/borja.sueiro/blockchain-restful-api/models"

type Blockchain interface {
	// Farm
	GetFarms() []models.Farm
	GetFarmById(string) (models.Farm, error)
	AddNewFarm(models.Farm) error
	UpdateFarm(models.Farm) error
	// Transport
	GetTransports() []models.Transport
	GetTransportById(string) (models.Transport, error)
	AddNewTransport(models.Transport) error
	UpdateTransport(models.Transport) error
	//Trace
	GetTraces() []models.Trace
	GetTraceById(string) (models.Trace, error)
	AddNewTrace(models.Trace) error
	UpdateTrace(models.Trace) error
}
