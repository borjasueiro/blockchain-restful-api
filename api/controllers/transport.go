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

func (ctrl *TransportController) GetTransports(c *gin.Context) {
	transports := ctrl.transportRepository.GetTransports()
	c.IndentedJSON(http.StatusOK, transports)
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
		c.Header("Location", common.TransportRoute+"/"+url.PathEscape(newTransport.ID))
	}
}

func (ctrl *TransportController) UpdateTransport(c *gin.Context) {
	var updatedTransport models.Transport
	if err := c.BindJSON(&updatedTransport); err != nil {
		return
	}
	if err := ctrl.transportRepository.UpdateTransport(updatedTransport); err != nil {
		if _, ok := err.(*blockchain.AssetNotFoundError); ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.Status(http.StatusNoContent)
	}

}
