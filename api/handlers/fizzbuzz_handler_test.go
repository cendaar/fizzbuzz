package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cendaar/fizzbuzz/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for FizzbuzzService
type MockFizzbuzzService struct {
	mock.Mock
}

func (m *MockFizzbuzzService) ComputeFizzbuzz(params models.FizzbuzzParams) string {
	args := m.Called(params)
	return args.String(0)
}

func (m *MockStatsRepo) LogRequest(params models.FizzbuzzParams) {
	m.Called(params)
}

func TestHandleFizzbuzz_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockFizzbuzzService := new(MockFizzbuzzService)
	mockStatsRepo := new(MockStatsRepo)
	fizzbuzzHandler := NewFizzbuzzHandler(mockFizzbuzzService, mockStatsRepo)
	r.POST("/fizzbuzz", fizzbuzzHandler.HandleFizzbuzz)

	// Invalid JSON request
	req, err := http.NewRequest(http.MethodPost, "/fizzbuzz", nil) // empty body
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "invalid request body"}`, w.Body.String())
}

func TestHandleFizzbuzz_ValidationError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	mockFizzbuzzService := new(MockFizzbuzzService)
	mockStatsRepo := new(MockStatsRepo)
	fizzbuzzHandler := NewFizzbuzzHandler(mockFizzbuzzService, mockStatsRepo)
	r.POST("/fizzbuzz", fizzbuzzHandler.HandleFizzbuzz)

	// Request with invalid params
	invalidParams := `{"int1": 0, "int2": 0, "limit": -1, "str1": "", "str2": ""}`
	req, err := http.NewRequest(http.MethodPost, "/fizzbuzz", bytes.NewBuffer([]byte(invalidParams)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestHandleFizzbuzz_ValidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Setup mock objects
	mockFizzbuzzService := new(MockFizzbuzzService)
	mockStatsRepo := new(MockStatsRepo)

	// Define valid params and expected response
	validParams := models.FizzbuzzParams{Int1: 3, Int2: 5, Limit: 15, Str1: "fizz", Str2: "buzz"}
	expectedResult := "1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz"

	// Set up mock behavior
	mockFizzbuzzService.On("ComputeFizzbuzz", validParams).Return(expectedResult)
	mockStatsRepo.On("LogRequest", validParams).Return()

	// Create the handler
	fizzbuzzHandler := NewFizzbuzzHandler(mockFizzbuzzService, mockStatsRepo)
	r.POST("/fizzbuzz", fizzbuzzHandler.HandleFizzbuzz)

	// Create a valid JSON request
	validRequest := `{"int1": 3, "int2": 5, "limit": 15, "str1": "fizz", "str2": "buzz"}`
	req, err := http.NewRequest(http.MethodPost, "/fizzbuzz", bytes.NewBuffer([]byte(validRequest)))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"result": "`+expectedResult+`"}`, w.Body.String())

	// Assert that the mock methods were called as expected
	mockFizzbuzzService.AssertExpectations(t)
	mockStatsRepo.AssertExpectations(t)
}
