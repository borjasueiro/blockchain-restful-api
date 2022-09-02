package api

import (
	"github.com/borja.sueiro/blockchain-restful-api/api/common"
	"github.com/borja.sueiro/blockchain-restful-api/api/controllers"
	"github.com/gin-gonic/gin"
)

type RestfulApplication struct {
	router        *gin.Engine
	transportCtrl *controllers.TransportController
	traceCtrl     *controllers.TraceController
}

func (app *RestfulApplication) Run(addr string) {
	app.router.Run(addr)
}

func setUpRoutes(router *gin.Engine,
	transportController *controllers.TransportController,
	traceController *controllers.TraceController) {

	// Transport
	router.GET(common.TransportRoute+"/:id", transportController.GetTransportById)
	router.POST(common.TransportRoute, transportController.AddNewTransport)
	router.POST(common.TransportRoute+"/:id/addfarm", transportController.AddFarmRecollectionToTransport)
	router.POST(common.TransportRoute+"/:id/popfarm", transportController.PopFarmRecollectionToTransport)
	// Trace
	router.GET(common.TraceRoute, traceController.GetTraces)
	router.GET(common.TraceRoute+"/:id", traceController.GetTraceById)
	router.POST(common.TraceRoute, traceController.AddNewTrace)
	router.POST(common.TraceRoute+"/:id/farm", traceController.AddFarmToTrace)
	router.POST(common.TraceRoute+"/:id/transvase", traceController.AddTransvaseToTrace)
}

func NewApp(transportCtrl *controllers.TransportController,
	traceCtrl *controllers.TraceController) *RestfulApplication {
	router := gin.Default()

	setUpRoutes(router, transportCtrl, traceCtrl)

	return &RestfulApplication{router, transportCtrl, traceCtrl}
}
