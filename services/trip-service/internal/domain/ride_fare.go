package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	UserID            string
	PackageSlug       string //ex: van , luxury, sedan
	TotalPriceInCents float64
	ExpiresAt         time.Time
}
