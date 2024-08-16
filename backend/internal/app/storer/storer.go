package storer

import (
	"context"
	"geomap/internal/app/models"
)

// CoordinateProvider
type CoordinateProvider interface {
	AllTracks(context.Context) ([]models.Track, error)
	LastPoints(ctx context.Context, tracks []models.Track) ([]models.Point, error)
	Route(ctx context.Context, track models.Track) ([]models.Position, error)
}
