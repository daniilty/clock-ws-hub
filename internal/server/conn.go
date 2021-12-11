package server

import (
	"net/http"
	"time"

	"context"

	"github.com/gorilla/websocket"
)

func (ws *WS) connHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		ws.logger.Debugw("Upgrade client.", "remoteAddr", r.RemoteAddr, "err", err)

		resp := getBadRequestWithMsgResponse("Failed to upgrade connection.")
		err = resp.writeJSON(w)
		if err != nil {
			ws.logger.Errorw("Write response.", "err", err)
		}

		return
	}

	ws.handleConn(r.Context(), conn)

	defer conn.Close()
}

// handleConn - send info about weather and portfolio after timeout
func (ws *WS) handleConn(ctx context.Context, conn *websocket.Conn) {
	const writeJSONErrMsg = "Write JSON."

	for {
		diff, err := ws.service.GetPortfolioDiff(ctx, ws.tinkoffAccountID)
		if err != nil {
			ws.logger.Errorw("GetPortfolioDiff", "err", err)
			msg := getErrorMessage("Failed to get portfolio diff.")

			err := conn.WriteJSON(msg)
			if err != nil {
				ws.logger.Errorw(writeJSONErrMsg, "err", err)
			}

			return
		}

		weather, err := ws.service.GetWeather(ctx)
		if err != nil {
			ws.logger.Errorw("GetWeather", "err", err)
			msg := getErrorMessage("Failed to get weather.")

			err := conn.WriteJSON(msg)
			if err != nil {
				ws.logger.Errorw(writeJSONErrMsg, "err", err)
			}

			return
		}

		msg := getDiffWeatherMessage(diff, weather)

		err = conn.WriteJSON(msg)
		if err != nil {
			ws.logger.Errorw(writeJSONErrMsg, "err", err)

			return
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(ws.timeout):
		}
	}
}
