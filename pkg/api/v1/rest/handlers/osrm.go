/*
 * telegram: @VasylNaumenko
 */

package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"routes-api/pkg/api/v1/errs"
	"routes-api/pkg/log"

	"routes-api/pkg/api/v1/rest/handlers/helpers/render"
	"routes-api/pkg/api/v1/services"
)

type Osrm struct {
	logger  log.Logger
	service services.Osrm
}

const dst = "dst"
const errNoSrc = "no src parameter"
const errNoDst = "at least one dst parameter must exists"

func NewOsrm(l log.Logger, service services.Osrm) Osrm {
	return Osrm{
		logger:  l,
		service: service,
	}
}

// processes routes query
func (h *Osrm) Process(w http.ResponseWriter, r *http.Request) {
	// extract the src parameter
	src := r.URL.Query().Get("src")
	if len(src) == 0 {

		render.Error(w, fmt.Errorf("%w%s", errs.ErrNotValid, errNoSrc))
		return
	}

	// store src as first parameter in payload
	payload := []string{
		src,
	}

	// extract dst parameter, at least one must exists
	raw := strings.Split(r.URL.RawQuery, "&")
	dstFound := false
	for _, s := range raw {
		params := strings.Split(s, "=")
		if len(params) != 2 {
			continue // skip wrong params
		}
		if params[0] == dst {
			payload = append(payload, params[1])
			dstFound = true
		}
	}
	if !dstFound {
		render.Error(w, fmt.Errorf("%w%s", errs.ErrNotValid, errNoDst))
		return
	}

	res, err := h.service.ProcessRequest(payload)
	if err != nil {
		render.Error(w, err)
		return
	}

	render.OK(w, res)
}
