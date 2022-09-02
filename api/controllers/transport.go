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

type TransportController struct {
	transportRepository *repo.TransportRepository
}

func NewTransportController(transportRepository *repo.TransportRepository) *TransportController {
	return &TransportController{transportRepository}
}

func (ctrl *TransportController) GetTransportById(c *gin.Context) {
	id := c.Param("id")
	if transport, err := ctrl.transportRepository.GetTransportById(id); err != nil {
		if _, ok := err.(*blockchain.AssetNotFoundError); ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.IndentedJSON(http.StatusOK, transport)
	}
}

func (ctrl *TransportController) AddNewTransport(c *gin.Context) {
	var newTransport models.Transport
	if err := c.BindJSON(&newTransport); err != nil {
		return // TODO return Error message.
	}

	if err := ctrl.transportRepository.AddNewTransport(newTransport); err != nil {
		if _, ok := err.(*blockchain.AssetAlreadyExistsError); ok {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.Status(http.StatusCreated)
		c.Header("Location", common.TransportRoute+"/"+url.PathEscape(newTransport.TransportID))
	}
}

func (ctrl *TransportController) AddFarmRecollectionToTransport(c *gin.Context) {
	transportId := c.Param("id")
	var farmRecollection models.FarmRecollection
	if err := c.BindJSON(&farmRecollection); err != nil {
		return
	}
	if err := ctrl.transportRepository.AddFarmRecollectionToTransport(transportId, farmRecollection); err != nil {
		if _, ok := err.(*blockchain.AssetNotFoundError); ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.Status(http.StatusNoContent)
	}
}

func (ctrl *TransportController) PopFarmRecollectionToTransport(c *gin.Context) {
	transportId := c.Param("id")
	farmName := c.PostForm("name")
	if farm, err := ctrl.transportRepository.PopFarmRecollectionToTransport(transportId, farmName); err != nil {
		if _, ok := err.(*blockchain.AssetNotFoundError); ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.IndentedJSON(http.StatusOK, farm)
	}
}
