package main

import (
	"flag"
	"net/http"
	delivery "symptomstracker/app/delivery"
	repository "symptomstracker/app/repository"
	usecase "symptomstracker/app/usecase"
	"symptomstracker/config"
	"time"
)

func main() {

	// DB configuration
	dbConfig := config.ConfigDB()

	// Router configuration
	router := config.ConfigRouter()

	rSymptomstrackerRepo := repository.NewRepository(dbConfig)

	uSymptomstrackerRepo := usecase.NewUsecase(rSymptomstrackerRepo)

	delivery.NewHTTPHandler(router, uSymptomstrackerRepo)

	s := &http.Server{
		Addr:         getPortNumber(),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	router.Logger.Fatal(router.StartServer(s))
}

func getPortNumber() string {
	port := flag.String("port", "8000", "Port number")
	flag.Parse()
	if port == nil || len(*port) == 0 {
		panic("Please enter port")
	}
	return ":" + *port
}
