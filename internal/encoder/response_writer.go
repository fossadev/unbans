package encoder

import (
	"encoding/json"
	"net/http"

	"github.com/fossadev/unbans/internal/logger"
)

type ResponseWriter interface {
	Success(interface{})
	Redirect(string)
	BadRequest(string)
	Unauthenticated(string)
	Forbidden(string)
	NotFound(string)
	Conflict(string)
	TooManyRequests(string)
	InternalServerError(string)

	Write([]byte) (int, error)
}

type responseWriter struct {
	http.ResponseWriter
	log logger.Logger
	req *http.Request
}

type errorResponse struct {
	Success    bool   `json:"success"`
	Status     int    `json:"status"`
	StatusText string `json:"status_text"`
	Error      string `json:"error"`
}

type successResponse struct {
	Success    bool        `json:"success"`
	Status     int         `json:"status"`
	StatusText string      `json:"status_text"`
	Response   interface{} `json:"response"`
}

const (
	locationHeader = "Location"
)

// NewResponseWriter returns a ResponseWriter for the given request.
// Errors from 5XX responses are logged to `log`.
func NewResponseWriter(w http.ResponseWriter, req *http.Request, log logger.Logger) ResponseWriter {
	return &responseWriter{
		ResponseWriter: w,
		log:            log,
		req:            req,
	}
}

func (w *responseWriter) writeJSON(status int, data interface{}) {
	if data == nil {
		w.WriteHeader(status)
		return
	}
	output, err := json.Marshal(data)
	if err != nil {
		w.log.Error("encoding err", err)
		status = http.StatusInternalServerError
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(output)
	if err != nil {
		w.log.Error("failed to write http response", err)
	}
}

func (w *responseWriter) writeError(status int, message string) {
	w.writeJSON(status, &errorResponse{
		Success:    false,
		Status:     status,
		StatusText: http.StatusText(status),
		Error:      message,
	})
}

func (w *responseWriter) Success(data interface{}) {
	w.writeJSON(http.StatusOK, data)
}

func (w *responseWriter) Redirect(loc string) {
	w.Header().Set(locationHeader, loc)
	w.WriteHeader(http.StatusFound)
}

func (w *responseWriter) BadRequest(err string) {
	w.writeError(http.StatusBadRequest, err)
}

func (w *responseWriter) Unauthenticated(err string) {
	w.writeError(http.StatusUnauthorized, err)
}

func (w *responseWriter) Forbidden(err string) {
	if err == "" {
		// fallback msg
		err = "You do not have permission to access this endpoint."
	}
	w.writeError(http.StatusForbidden, err)
}

func (w *responseWriter) NotFound(err string) {
	if err == "" {
		// fallback msg
		err = "404 page not found"
	}
	w.writeError(http.StatusNotFound, err)
}

func (w *responseWriter) Conflict(err string) {
	w.writeError(http.StatusConflict, err)
}

func (w *responseWriter) TooManyRequests(err string) {
	w.writeError(http.StatusTooManyRequests, err)
}

func (w *responseWriter) InternalServerError(err string) {
	w.writeError(http.StatusInternalServerError, err)
}
