package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cendaar/fizzbuzz/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for StatisticsRepository
type MockStatsRepo struct {
	mock.Mock
}

func (m *MockStatsRepo) GetMostFrequentRequest() (*models.FizzbuzzParams, int) {
	args := m.Called()
	return args.Get(0).(*models.FizzbuzzParams), args.Int(1)
}

func TestHandleStats_Success(t *testing.T) {
	// Initialize Gin
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Create a mock StatisticsRepository
	mockRepo := new(MockStatsRepo)
	mockRepo.On("GetMostFrequentRequest").Return(
		&models.FizzbuzzParams{
			Int1:  2,
			Int2:  3,
			Limit: 10,
			Str1:  "Hello",
			Str2:  "World",
		},
		3,
	)

	// Create StatsHandler with mock repository
	statsHandler := NewStatsHandler(mockRepo)
	r.GET("/stats", statsHandler.HandleStats)

	// Create a test request
	req, err := http.NewRequest(http.MethodGet, "/stats", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"result":{"hits":3,"most_common_request":{"int1":2,"int2":3,"limit":10,"str1":"Hello","str2":"World"}}}`, w.Body.String())

	// Assert that the mock was called
	mockRepo.AssertExpectations(t)
}

func TestHandleStats_NoStats(t *testing.T) {
	// Initialize Gin
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Create a mock StatisticsRepository
	mockRepo := new(MockStatsRepo)
	mockRepo.On("GetMostFrequentRequest").Return(
		&models.FizzbuzzParams{},
		0,
	)

	// Create StatsHandler with mock repository
	statsHandler := NewStatsHandler(mockRepo)
	r.GET("/stats", statsHandler.HandleStats)

	// Create a test request
	req, err := http.NewRequest(http.MethodGet, "/stats", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusNoContent, w.Code)

	// Assert that the mock was called
	mockRepo.AssertExpectations(t)
}
