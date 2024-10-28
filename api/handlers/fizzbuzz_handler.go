package handlers

import (
	"net/http"

	"github.com/cendaar/fizzbuzz/models"
	"github.com/cendaar/fizzbuzz/repository"
	"github.com/cendaar/fizzbuzz/services"
	"github.com/gin-gonic/gin"
)

type FizzbuzzHandler struct {
	fizzbuzzService services.FizzbuzzServiceI
	statsRepo       repository.StatisticsRepositoryI
}

func NewFizzbuzzHandler(fbs services.FizzbuzzServiceI, statsRepo repository.StatisticsRepositoryI) *FizzbuzzHandler {
	return &FizzbuzzHandler{
		fizzbuzzService: fbs,
		statsRepo:       statsRepo,
	}
}

func (fbh *FizzbuzzHandler) HandleFizzbuzz(c *gin.Context) {
	var params models.FizzbuzzParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := params.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fbh.statsRepo.LogRequest(params)

	result := fbh.fizzbuzzService.ComputeFizzbuzz(params)
	c.JSON(http.StatusOK, gin.H{"result": result})
}
