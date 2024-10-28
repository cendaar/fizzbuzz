package handlers

import (
	"net/http"

	"github.com/cendaar/fizzbuzz/repository"
	"github.com/gin-gonic/gin"
)

type StatsHandler struct {
	statsRepo repository.StatisticsRepositoryI
}

func NewStatsHandler(statsRepo repository.StatisticsRepositoryI) *StatsHandler {
	return &StatsHandler{
		statsRepo: statsRepo,
	}
}

func (s *StatsHandler) HandleStats(c *gin.Context) {
	mostCommonRequest, maxHits := s.statsRepo.GetMostFrequentRequest()

	if maxHits == 0 {
		c.JSON(http.StatusNoContent, nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": map[string]interface{}{
		"most_common_request": mostCommonRequest,
		"hits":                maxHits,
	}})
}
