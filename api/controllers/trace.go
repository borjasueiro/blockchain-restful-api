package controllers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/borja.sueiro/blockchain-restful-api/api/common"
	repo "github.com/borja.sueiro/blockchain-restful-api/api/repositories"
	"github.com/borja.sueiro/blockchain-restful-api/blockchain"
	"github.com/borja.sueiro/blockchain-restful-api/models"
	"github.com/gin-gonic/gin"
)

type TraceController struct {
	traceRepository *repo.TraceRepository
}

func NewTraceController(traceRepository *repo.TraceRepository) *TraceController {
	return &TraceController{traceRepository}
}

func (ctrl *TraceController) GetTraces(c *gin.Context) {
	traces := ctrl.traceRepository.GetTraces()
	c.IndentedJSON(http.StatusOK, traces)
}

func (ctrl *TraceController) GetTraceById(c *gin.Context) {
	id := c.Param("id")
	if trace, err := ctrl.traceRepository.GetTraceById(id); err != nil {
		if _, ok := err.(*blockchain.AssetNotFoundError); ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.IndentedJSON(http.StatusOK, trace)
	}
}

func (ctrl *TraceController) AddNewTrace(c *gin.Context) {
	var newFarm models.FarmRecollection
	if err := c.BindJSON(&newFarm); err != nil {
		return // TODO return Error message.
	}
	id, err := ctrl.traceRepository.AddNewTrace(newFarm)
	if err != nil {
		if _, ok := err.(*blockchain.AssetAlreadyExistsError); ok {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.Status(http.StatusCreated)
		c.Header("Location", common.TraceRoute+"/"+url.PathEscape(id))
	}
}

func (ctrl *TraceController) AddFarmToTrace(c *gin.Context) {
	var farm models.FarmRecollection
	id := c.Param("id")
	if err := c.BindJSON(&farm); err != nil {
		return
	}
	if err := ctrl.traceRepository.AddFarmToTrace(id, farm); err != nil {
		if _, ok := err.(*blockchain.AssetNotFoundError); ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.Status(http.StatusNoContent)
	}

}

func (ctrl *TraceController) AddTransvaseToTrace(c *gin.Context) {
	var transvase models.Transvase
	id := c.Param("id")
	if err := c.BindJSON(&transvase); err != nil {
		return
	}
	if err := ctrl.traceRepository.AddTransvaseToTrace(id, transvase); err != nil {
		if _, ok := err.(*blockchain.AssetNotFoundError); ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.Status(http.StatusNoContent)
	}

}
