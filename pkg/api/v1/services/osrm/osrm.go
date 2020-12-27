/*
 * telegram: @VasylNaumenko
 */

package osrm

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"

	"routes-api/pkg/api/v1/models"
	"routes-api/pkg/api/v1/rest/handlers/helpers"
	"routes-api/pkg/api/v1/services"
	"routes-api/pkg/config"
	"routes-api/pkg/log"
)

// type Service implements services Osrm interface
type serviceOsrm struct {
	cfg config.RestVendors
	log log.Logger
}

const params = "overview=false&generate_hints=false"
const ok = "Ok"

func (s serviceOsrm) ProcessRequest(payload []string) (models.Response, error) {
	var rawResp models.OsrmResponse
	var resp models.Response

	// making of a request url
	url := fmt.Sprintf("%s%s?%s",
		s.cfg.Osrm,
		strings.Join(payload, ";"),
		params,
	)

	// sending request and getting response or error
	err := helpers.GetUrlResponse(url, &rawResp)
	if err != nil {
		return resp, err
	}

	// got an error
	if rawResp.Code != ok {
		return resp, errors.New(fmt.Sprint(rawResp.Code, rawResp.Message))
	}

	// parsing vendor response into service response model
	waypoints := rawResp.Waypoints
	resp.Source = printLocation(waypoints[0].Location)
	for i, route := range rawResp.Routes[0].Legs {
		dest := models.Route{
			Distance: route.Distance,
			Duration: route.Duration,
		}

		if len(waypoints) > i+1 && len(waypoints[i].Location) == 2 {
			dest.Destination = fmt.Sprintf("%f,%f", waypoints[i].Location[0], waypoints[i].Location[1])
		}

		resp.Routes = append(resp.Routes, dest)
	}

	// sorting results
	sort.Sort(resp.Routes)

	return resp, nil
}

func printLocation(l []float32) string {
	return fmt.Sprintf("%f,%f", l[0], l[1])
}

func New(c config.RestVendors, l log.Logger) services.Osrm {
	return &serviceOsrm{
		cfg: c,
		log: l,
	}
}
