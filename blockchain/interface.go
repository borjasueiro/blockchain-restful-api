package blockchain

import (
	"github.com/borja.sueiro/blockchain-restful-api/models"
)

type Blockchain interface {
	// Transport
	GetTransportById(string) (models.Transport, error)
	AddNewTransport(models.Transport) error
	AddFarmRecollectionToTransport(string, models.FarmRecollection) error
	PopFarmRecollectionToTransport(string, string) (models.FarmRecollection, error)
	//Trace
	GetTraces() []models.Trace
	GetTraceById(string) (models.Trace, error)
	AddNewTrace(models.FarmRecollection) (string, error)
	AddFarmToTrace(string, models.FarmRecollection) error
	AddTransvaseToTrace(string, models.Transvase) error
}
