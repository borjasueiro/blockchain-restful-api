package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/borja.sueiro/blockchain-restful-api/models"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type HyperledgerApp struct {
	contract *gateway.Contract
}

func NewHyperledgerApp() *HyperledgerApp {
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Failed to populate wallet contents: %v", err)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	contract := network.GetContract("basic")
	log.Println("--> Submit Transaction: InitLedger, function creates the initial set of assets on the ledger")
	result, err := contract.SubmitTransaction("InitLedger")
	if err != nil {
		log.Fatalf("Failed to Submit transaction: %v", err)
	}
	log.Println(string(result))

	// log.Println("--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")
	// result, err = contract.EvaluateTransaction("GetAllAssets")
	// if err != nil {
	// 	log.Fatalf("Failed to evaluate transaction: %v", err)
	// }
	// log.Println(string(result))
	return &HyperledgerApp{contract}
}

func populateWallet(wallet *gateway.Wallet) error {
	log.Println("============ Populating wallet ============")
	credPath := filepath.Join(
		"..",
		"fabric-samples",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	return wallet.Put("appUser", identity)
}

func (app *HyperledgerApp) GetFarms() []models.Farm {
	result, err := app.contract.EvaluateTransaction("GetAllFarms")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	var farms []models.Farm
	if err = json.Unmarshal(result, &farms); err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	return farms
}

func (app *HyperledgerApp) GetFarmById(id string) (models.Farm, error) {
	result, err := app.contract.EvaluateTransaction("ReadFarm", id)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return models.Farm{}, err
	}
	var farm models.Farm
	if err = json.Unmarshal(result, &farm); err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return models.Farm{}, err
	}
	return farm, nil
}

func (app *HyperledgerApp) AddNewFarm(farm models.Farm) error {
	_, err := app.contract.SubmitTransaction("CreateFarm", farm.ID, farm.Location)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return err
	}
	return nil
}

func (app *HyperledgerApp) UpdateFarm(farm models.Farm) error {
	_, err := app.contract.SubmitTransaction("UpdateFarm", farm.ID, farm.Location)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return err
	}

	return nil
}

// Transport
func (app *HyperledgerApp) GetTransports() []models.Transport {
	result, err := app.contract.EvaluateTransaction("GetAllTransports")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	var transports []models.Transport
	if err = json.Unmarshal(result, &transports); err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	return transports
}

func (app *HyperledgerApp) GetTransportById(id string) (models.Transport, error) {
	result, err := app.contract.EvaluateTransaction("ReadTransport", id)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return models.Transport{}, err
	}
	var transport models.Transport
	if err = json.Unmarshal(result, &transport); err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return models.Transport{}, err
	}
	return transport, nil
}

func (app *HyperledgerApp) AddNewTransport(transport models.Transport) error {
	_, err := app.contract.SubmitTransaction("CreateTransport", transport.TransportID, transport.SiloID, transport.Date)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return err
	}
	return nil
}

func (app *HyperledgerApp) UpdateTransport(transport models.Transport) error {
	_, err := app.contract.SubmitTransaction("UpdateTransport", transport.TransportID, transport.SiloID, transport.Date)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return err
	}

	return nil
}

// Traces
func (app *HyperledgerApp) GetTraces() []models.Trace {
	result, err := app.contract.EvaluateTransaction("GetAllTraces")
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	var traces []models.Trace
	if err = json.Unmarshal(result, &traces); err != nil {
		log.Fatalf("Failed to evaluate transaction: %v", err)
	}
	return traces
}

func (app *HyperledgerApp) GetTraceById(id string) (models.Trace, error) {
	result, err := app.contract.EvaluateTransaction("ReadTrace", id)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return models.Trace{}, err
	}
	var trace models.Trace
	if err = json.Unmarshal(result, &trace); err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return models.Trace{}, err
	}
	return trace, nil
}

func (app *HyperledgerApp) AddNewTrace(id string, farm models.Farm) error {
	_, err := app.contract.SubmitTransaction("CreateTrace", id, farm.ID, farm.Location, farm.Date, farm.TransportId, farm.Temperature)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return err
	}
	return nil
}

func (app *HyperledgerApp) AddFarmToTrace(id string, farm models.Farm) error {
	_, err := app.contract.SubmitTransaction("AddFarmToTrace", id, farm.ID, farm.Location, farm.Date, farm.TransportId, farm.Temperature)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return err
	}
	return nil
}

func (app *HyperledgerApp) AddTransvaseToTrace(id string, transvase models.Transvase) error {
	_, err := app.contract.SubmitTransaction("AddTransvaseToTrace", id, transvase.SrcSiloID, transvase.DstSiloID, transvase.Date)
	if err != nil {
		log.Fatalf("Failed to evaluate transaction: %v\n", err)
		return err
	}

	return nil
}
