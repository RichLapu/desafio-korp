package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Métrica 1: Volume de requisições[cite: 1]
	requestVolume = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total de requisições HTTP para o endpoint /projeto-korp",
		},
	)
	// Métrica 2: Disponibilidade do serviço[cite: 1]
	serviceAvailability = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "service_availability",
			Help: "Status de disponibilidade do serviço (1 = UP, 0 = DOWN)",
		},
	)
)

func init() {
	// Registra as métricas no Prometheus
	prometheus.MustRegister(requestVolume)
	prometheus.MustRegister(serviceAvailability)
	
	// Define o serviço como disponível (1) logo na inicialização
	serviceAvailability.Set(1)
}

type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	// Incrementa a métrica de volume de requisições toda vez que a rota é chamada
	requestVolume.Inc()

	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Rota principal do projeto
	http.HandleFunc("/projeto-korp", projetoKorpHandler)
	
	// Rota padrão exigida para expor as métricas para o Prometheus[cite: 1]
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}