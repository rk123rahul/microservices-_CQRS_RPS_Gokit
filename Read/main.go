package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	// "github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	// kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	// stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	db := GetDBconn()

	r := mux.NewRouter()

	var svc AccountService
	svc = accountservice{}
	{
		repository, err := NewRepo(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = NewService(repository, logger)
	}
	// svc = loggingMiddleware{logger, svc}
	// svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	
	GetByCustomerIdHandler := httptransport.NewServer(
		makeGetCustomerByIdEndpoint(svc),
		decodeGetCustomerByIdRequest,
		encodeResponse,
	)
	GetAllCustomersHandler := httptransport.NewServer(
		makeGetAllCustomersEndpoint(svc),
		decodeGetAllCustomersRequest,
		encodeResponse,
	)
	

	http.Handle("/", r)
	
	
	r.Handle("/account/getAll", GetAllCustomersHandler).Methods("GET")
	r.Handle("/account/{customerid}", GetByCustomerIdHandler).Methods("GET")
	

	// http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
