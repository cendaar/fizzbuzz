package repository

import (
	"sync"

	"github.com/cendaar/fizzbuzz/models"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type StatisticsRepositoryI interface {
	GetMostFrequentRequest() (*models.FizzbuzzParams, int)
	LogRequest(params models.FizzbuzzParams)
}

// StatisticsRepository stores and tracks FizzBuzz request statistics.
type StatisticsRepository struct {
	requests       map[models.FizzbuzzParams]int
	mu             sync.Mutex
	fizzBuzzCalls  prometheus.Counter
	mostUsedParams *prometheus.CounterVec
	mostFrequent   struct {
		params *models.FizzbuzzParams
		hits   int
	}
}

// NewStatisticsRepository initializes a new StatisticsRepository.
func NewStatisticsRepository() *StatisticsRepository {
	return &StatisticsRepository{
		requests: make(map[models.FizzbuzzParams]int),
		fizzBuzzCalls: promauto.NewCounter(prometheus.CounterOpts{
			Name: "fizzbuzz_calls_total",
			Help: "Total number of FizzBuzz calls.",
		}),
		mostUsedParams: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "fizzbuzz_request_params_total",
			Help: "Total number of FizzBuzz requests.",
		}, []string{"params"}),
	}
}

// LogRequest records a FizzBuzz request and updates statistics.
func (repo *StatisticsRepository) LogRequest(params models.FizzbuzzParams) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.requests[params]++

	// Update Prometheus metrics
	repo.fizzBuzzCalls.Inc()
	repo.mostUsedParams.WithLabelValues(
		params.String(),
	).Inc()

	// Update most frequent request cache if necessary
	if repo.requests[params] > repo.mostFrequent.hits {
		repo.mostFrequent.params = &params
		repo.mostFrequent.hits = repo.requests[params]
	}
}

// GetMostFrequentRequest returns the most frequently used parameters and hit count.
func (repo *StatisticsRepository) GetMostFrequentRequest() (*models.FizzbuzzParams, int) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	return repo.mostFrequent.params, repo.mostFrequent.hits
}
