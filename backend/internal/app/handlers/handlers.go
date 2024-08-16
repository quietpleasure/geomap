package handlers

import (
	"context"
	"fmt"
	"geomap/internal/app/helper"
	"geomap/internal/app/logger"
	"geomap/internal/app/models"
	"geomap/internal/app/storer"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	log  *logger.Logger
	help *helper.Helper
}

func New(log *logger.Logger, store storer.CoordinateProvider) *Handler {
	l := &logger.Logger{Logger: log.Named("[HANDLER]")}
	l.Debug("initialized package")
	return &Handler{
		log:  l,
		help: helper.New(store),
	}
}

func (h Handler) AllTracks() gin.HandlerFunc {
	return func(c *gin.Context) {

		// c.Error(fmt.Errorf("my custom internal error #1"))
		// c.Error(fmt.Errorf("my internal error #2"))
		// c.Errors = append(c.Errors, )
		// id := c.Writer.Header().Get("X-Correlation-ID") можно получать из хэдера
		id := c.GetString("X-Correlation-ID")

		tracks, err := h.help.AllTracks(context.Background())
		if err != nil {
			c.Error(fmt.Errorf("get all tracks: %s", err))
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		// render.Index(c, status, tracks)
		h.log.Debug("data send response", zap.String("correlation-id", id), zap.Any("tracks", tracks))
		c.JSON(http.StatusOK, tracks)
	}
}

func (h Handler) ShowTracks() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString("X-Correlation-ID")

		var tracks []models.Track
		if err := c.Bind(&tracks); err != nil {
			// h.log.Error("binding data", zap.Error(err))
			c.Error(fmt.Errorf("binding data: %s", err))
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		h.log.Debug("data from request", zap.String("correlation-id", id), zap.Any("tracks", tracks))
		points, err := h.help.LastCoordinate(context.Background(), tracks)
		if err != nil {
			c.Error(fmt.Errorf("get last points: %s", err))
			// status = http.StatusInternalServerError
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		h.log.Debug("data send response", zap.String("correlation-id", id), zap.Any("points", points))
		c.JSON(http.StatusOK, points)
	}
}

func (h Handler) ShowTrack() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetString("X-Correlation-ID")

		var track models.Track
		if err := c.Bind(&track); err != nil {
			// h.log.Error("binding data", zap.Error(err))
			c.Error(fmt.Errorf("binding data: %s", err))
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		h.log.Debug("data from request", zap.String("correlation-id", id), zap.Any("track", track))
		points, err := h.help.Route(context.Background(), track)
		if err != nil {
			c.Error(fmt.Errorf("get last points: %s", err))
			// status = http.StatusInternalServerError
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
		smooth := h.help.RDPAlgorithm(0.00799, points)
		h.log.Debug("data send response", zap.String("correlation-id", id), zap.Any("points", smooth))
		c.JSON(http.StatusOK, smooth)
	}
}
