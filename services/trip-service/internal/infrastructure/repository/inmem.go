package repository

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
)

type InMemTripRepository struct {
	trips    map[string]*domain.TripModel
	RideFare map[string]*domain.RideFareModel
}

func NewInMemTripRepository() *InMemTripRepository {
	return &InMemTripRepository{
		trips:    make(map[string]*domain.TripModel),
		RideFare: make(map[string]*domain.RideFareModel),
	}
}

func (r *InMemTripRepository) CreateTrip(ctx context.Context, trip *domain.TripModel) (*domain.TripModel, error) {
	r.trips[trip.ID.Hex()] = trip
	return trip, nil
}
