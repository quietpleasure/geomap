package postgres

import (
	"context"
	"fmt"

	"geomap/internal/app/models"

	"github.com/jackc/pgx/v5"
)

func (p *Postgres) AllTracks(ctx context.Context) ([]models.Track, error) {
	rows, err := p.Query(ctx, `select uniqueid, model, contact from tc_devices`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tracks []models.Track
	for rows.Next() {
		var t models.Track
		if err := rows.Scan(
			&t.ID,
			&t.Number,
			&t.Driver,
		); err != nil {
			return nil, err
		}
		tracks = append(tracks, t)
	}
	return tracks, nil
}

func (p *Postgres) LastPoints(ctx context.Context, tracks []models.Track) ([]models.Point, error) {
	var points []models.Point
	for label, track := range tracks {
		point := models.Point{
			Lable:    fmt.Sprintf("%d", label+1),
			Position: models.Position{},
		}
		var (
			speed, course float64
			attributes    string
		)

		if err := p.QueryRow(ctx,
			`SELECT tp.latitude , tp.longitude ,tp.speed ,tp.course ,tp."attributes"  
			FROM tc_positions tp 
			JOIN tc_devices td ON td.id = tp.deviceid 
			WHERE td.uniqueid = $1 ORDER BY tp.fixtime DESC LIMIT 1`,
			track.ID,
		).Scan(
			&point.Position.Lat,
			&point.Position.Lng,
			&speed,
			&course,
			&attributes,
		); err != nil && err != pgx.ErrNoRows {
			return nil, err
		}
		point.Title = fmt.Sprintf("%s - %s", track.Number, track.Driver)
		points = append(points, point)
	}
	return points, nil
}

func (p *Postgres) Route(ctx context.Context, track models.Track) ([]models.Position, error) {
	rows, err := p.Query(ctx,
		`SELECT DISTINCT ON (tp.latitude , tp.longitude) tp.latitude , tp.longitude, tp.id  FROM tc_positions tp 
		JOIN tc_devices td ON td.id = tp.deviceid 
		WHERE td.uniqueid = $1 
		AND tp.fixtime >= (
							SELECT max(tp2.fixtime) FROM tc_positions tp2 JOIN tc_devices td2 ON td2.id = tp2.deviceid WHERE td2.uniqueid = $1
							)-interval '3 day'`,
		track.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var points []models.Position
	for rows.Next() {
		var (
			id int64
			pnt models.Position
		)
		if err := rows.Scan(
			&pnt.Lat,
			&pnt.Lng,
			&id,
		); err != nil {
			return nil, err
		}

		points = append(points, pnt)
	}
	return points, nil
}
