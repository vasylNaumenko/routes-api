/*
 * telegram: @VasylNaumenko
 */

package services

import (
	"routes-api/pkg/api/v1/models"
)

type Osrm interface {
	ProcessRequest(payload []string) (models.Response, error)
}
