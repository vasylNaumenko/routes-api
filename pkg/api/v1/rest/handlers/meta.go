/*
 * telegram: @VasylNaumenko
 */

package handlers

import (
	"net/http"

	"routes-api/pkg/log"

	"routes-api/pkg/api/v1/rest/handlers/helpers/render"
)

type Meta struct {
	logger log.Logger
}

func NewMeta(l log.Logger) Meta {
	return Meta{
		logger: l,
	}
}

// Status used for service health check
func (h Meta) Status(w http.ResponseWriter, _ *http.Request) {
	render.OK(w,
		struct {
			Status string
		}{"ok"},
	)
}
