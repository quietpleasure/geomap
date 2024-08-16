package helper

import (
	"context"
	"geomap/internal/app/models"
	"geomap/internal/app/storer"

	rdp "github.com/calvinfeng/rdp-path-simplification"
)

type Helper struct {
	store storer.CoordinateProvider
}

func New(store storer.CoordinateProvider) *Helper {
	return &Helper{store: store}
}

func (h *Helper) LastCoordinate(ctx context.Context, tracks []models.Track) ([]models.Point, error) {
	return h.store.LastPoints(ctx, tracks)
}

func (h *Helper) AllTracks(ctx context.Context) ([]models.Track, error) {
	return h.store.AllTracks(ctx)
}

func (h *Helper) Route(ctx context.Context, track models.Track) ([]models.Position, error) {
	return h.store.Route(ctx, track)
}

func (h *Helper) RDPAlgorithm(threshold float64, input []models.Position) []models.Position {
	points := make([]rdp.Point, 0, len(input))
	for _, p := range input {
		points = append(points, rdp.Point{
			X: p.Lat,
			Y: p.Lng,
		})
	}
	alg := rdp.SimplifyPath(points, threshold)
	out := make([]models.Position, 0, len(alg))
	for _, p := range alg {
		out = append(out, models.Position{
			Lat: p.X,
			Lng: p.Y,
		})
	}
	return out
}
