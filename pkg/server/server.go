package server

import (
	"github.com/emilstorgaardandersen/stockDataApi/pkg/handlers/reading"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Name   string
	router *mux.Router
}

func New(name string) (*Server, error) {
	r := mux.NewRouter()

	s := r.PathPrefix("/api/v1").Subrouter()
	s.Path("/stockData/{id}").Handler(reading.GetData()) // /api/v1/stockData/AAPL
	s.Path("/stocksData").Handler(reading.GetMultiData()) // /api/v1/stocksData?stocks=TSLA&stocks=AAPL
	s.Path("/myPortfolio").Handler(reading.GetMyPortfolio()) // /api/v1/myPortfolio

	return &Server{Name: name, router: r}, nil
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.router)
}