package handler

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ControllerHandler func(w http.ResponseWriter, r *http.Request) (int, any)

func ToLoggedHandlerFunc(controllerHandler ControllerHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := log.With().
			Str("path", r.URL.String()).
			Str("method", r.Method).Logger()
		logger.Info().Msg("request started")

		status, body := controllerHandler(w, r)
		w.WriteHeader(status)
		if body != nil {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(body)
			if err != nil {
				logger.Err(err).Send()
			}
		}

		logger.Info().
			Int("status", status).
			Msg("request completed")
	}
}
