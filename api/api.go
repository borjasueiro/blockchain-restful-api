package api

import (
	"github.com/borja.sueiro/blockchain-restful-api/api/common"
	"github.com/borja.sueiro/blockchain-restful-api/api/controllers"
	"github.com/gin-gonic/gin"
)

type RestfulApplication struct {
	router         *gin.Engine
	farmController *controllers.FarmController
	transportCtrl  *controllers.TransportController
	traceCtrl      *controllers.TraceController
}

func (app *RestfulApplication) Run(addr string) {
	app.router.Run(addr)
}

func setUpRoutes(router *gin.Engine, farmController *controllers.FarmController,
	transportController *controllers.TransportController,
	traceController *controllers.TraceController) {
	// Farm
	router.GET(common.FarmRoute, farmController.GetFarms)
	router.GET(common.FarmRoute+"/:id", farmController.GetFarmById)
	router.POST(common.FarmRoute, farmController.AddNewFarm)
	router.PUT(common.FarmRoute+"/:id", farmController.UpdateFarm)
	// Transport
	router.GET(common.TransportRoute, transportController.GetTransports)
	router.GET(common.TransportRoute+"/:id", transportController.GetTransportById)
	router.POST(common.TransportRoute, transportController.AddNewTransport)
	router.PUT(common.TransportRoute+"/:id", transportController.UpdateTransport)
	// Trace
	router.GET(common.TraceRoute, traceController.GetTraces)
	router.GET(common.TraceRoute+"/:id", traceController.GetTraceById)
	router.POST(common.TraceRoute, traceController.AddNewTrace)
	router.PUT(common.TraceRoute+"/:id", traceController.UpdateTrace)
}

func NewApp(farmCtrl *controllers.FarmController,
	transportCtrl *controllers.TransportController,
	traceCtrl *controllers.TraceController) *RestfulApplication {
	router := gin.Default()

	setUpRoutes(router, farmCtrl, transportCtrl, traceCtrl)

	return &RestfulApplication{router, farmCtrl, transportCtrl, traceCtrl}
}
