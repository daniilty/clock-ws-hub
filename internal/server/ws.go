package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/daniilty/clock-ws-hub/internal/core"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type WS struct {
	httpServer       *http.Server
	logger           *zap.SugaredLogger
	service          core.Service
	upgrader         *websocket.Upgrader
	timeout          time.Duration
	tinkoffAccountID string
}

func NewWS(service core.Service, logger *zap.SugaredLogger, timeout time.Duration, addr string, accountID string) *WS {
	ws := &WS{
		logger:           logger,
		service:          service,
		upgrader:         &websocket.Upgrader{},
		timeout:          timeout,
		tinkoffAccountID: accountID,
	}

	// disable origin check
	ws.upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	r := mux.NewRouter()
	ws.setRoutes(r)

	ws.httpServer = &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return ws
}

func (h *WS) Run(ctx context.Context) {
	h.logger.Infow("HTTP server starting.", "addr", h.httpServer.Addr)

	go func() {
		err := h.httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}

			h.logger.Errorw("Listen and serve HTTP", "addr", h.httpServer.Addr, "err", err)
		}
	}()

	<-ctx.Done()

	h.logger.Info("Graceful server shutdown.")
	h.httpServer.Shutdown(context.Background())
}
