package api

import (
	"github.com/cendaar/fizzbuzz/api/handlers"
	"github.com/cendaar/fizzbuzz/repository"
	"github.com/cendaar/fizzbuzz/services"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	router          *gin.Engine
	fizzbuzzHandler *handlers.FizzbuzzHandler
	statsHandler    *handlers.StatsHandler
}

func NewServer() *Server {
	fizzBuzzService := services.NewFizzbuzzService()
	statsRepo := repository.NewStatisticsRepository()

	fizzbuzzHandler := handlers.NewFizzbuzzHandler(fizzBuzzService, statsRepo)
	statsHandler := handlers.NewStatsHandler(statsRepo)

	return &Server{
		router:          gin.New(),
		fizzbuzzHandler: fizzbuzzHandler,
		statsHandler:    statsHandler,
	}
}

func (s *Server) Start() error {
	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())

	s.registerRoutes()

	return s.router.Run()
}

func (s *Server) registerRoutes() {
	s.router.GET("/fizzbuzz", s.fizzbuzzHandler.HandleFizzbuzz)
	s.router.GET("/stats", s.statsHandler.HandleStats)
	s.router.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
