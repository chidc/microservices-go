package main

import (
	"log"
	"net/http"
	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
)

func main() {
	//ctx := context.Background()
	inmemRepo := repository.NewInMemTripRepository()
	svc := service.NewTripService(inmemRepo)
	mux := http.NewServeMux()

	httphandler := h.HttpHandler{Service: svc}
	// t, err := tripService.CreateTrip(ctx, &domain.RideFareModel{
	// 	UserID:      "42",
	// 	PackageSlug: "sedan",
	// })
	mux.HandleFunc("POST /preview", httphandler.HandleTripPreview)
	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println(t)
	// for {
	// 	time.Sleep(time.Second)
	// }
}
