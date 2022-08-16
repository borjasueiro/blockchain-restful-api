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

type FarmController struct {
	farmRepository *repo.FarmRepository
}

func NewFarmController(farmRepository *repo.FarmRepository) *FarmController {
	return &FarmController{farmRepository}
}

func (ctrl *FarmController) GetFarms(c *gin.Context) {
	farms := ctrl.farmRepository.GetFarms()
	c.IndentedJSON(http.StatusOK, farms)
}

func (ctrl *FarmController) GetFarmById(c *gin.Context) {
	id := c.Param("id")
	if farm, err := ctrl.farmRepository.GetFarmById(id); err != nil {
		if _, ok := err.(*blockchain.AssetNotFoundError); ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.IndentedJSON(http.StatusOK, farm)
	}
}

func (ctrl *FarmController) AddNewFarm(c *gin.Context) {
	var newFarm models.Farm
	if err := c.BindJSON(&newFarm); err != nil {
		return // TODO return Error message.
	}

	if err := ctrl.farmRepository.AddNewFarm(newFarm); err != nil {
		if _, ok := err.(*blockchain.AssetAlreadyExistsError); ok {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.Status(http.StatusCreated)
		c.Header("Location", common.FarmRoute+"/"+url.PathEscape(newFarm.ID))
	}
}

func (ctrl *FarmController) UpdateFarm(c *gin.Context) {
	var updatedFarm models.Farm
	if err := c.BindJSON(&updatedFarm); err != nil {
		return
	}
	if err := ctrl.farmRepository.UpdateFarm(updatedFarm); err != nil {
		if _, ok := err.(*blockchain.AssetNotFoundError); ok {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprint(err)})
		} else {
			c.Status(http.StatusInternalServerError)
		}
	} else {
		c.Status(http.StatusNoContent)
	}

}
