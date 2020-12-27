/*
 * telegram: @VasylNaumenko
 */

package rest

import (
	"github.com/go-chi/chi"
)

// REST API routes
func (a *Api) registerRoutes(r chi.Router) {
	r.Route("/", func(r chi.Router) {
		r.Get("/status", a.hMeta.Status)
		r.Get("/routes", a.hOsrm.Process)
	})
}
