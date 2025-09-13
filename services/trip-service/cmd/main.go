package main

import (
	"context"
	"log"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	ctx := context.Background()
	inmemRepo := repository.NewInMemTripRepository()
	tripService := service.NewTripService(inmemRepo)

	t, err := tripService.CreateTrip(ctx, &domain.RideFareModel{
		UserID:      "42",
		PackageSlug: "sedan",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(t)
	for {
		time.Sleep(time.Second)
	}
}
