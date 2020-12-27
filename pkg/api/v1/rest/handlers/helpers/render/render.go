/*
 * telegram: @VasylNaumenko
 */

package render

import (
	"encoding/json"
	"errors"
	"net/http"

	"routes-api/pkg/api/v1/errs"
)

const (
	statusError = "error"
)

type BodyError struct {
	Status  string `json:"status" default:"error"`
	Message string `json:"message" example:"something went wrong"`
}

func Error(w http.ResponseWriter, err error) {
	respondError(w, resolveCode(err), err)
}

func OK(w http.ResponseWriter, data interface{}) {
	respondSuccess(w, http.StatusOK, data)
}

func resolveCode(err error) int {
	if errors.Is(err, errs.ErrNotFound) {
		return http.StatusNotFound
	}
	if errors.Is(err, errs.ErrNotValid) {
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}

func respondError(w http.ResponseWriter, code int, err error) {
	var r BodyError
	r.Status = statusError
	r.Message = err.Error()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(r)
}

func respondSuccess(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
}
