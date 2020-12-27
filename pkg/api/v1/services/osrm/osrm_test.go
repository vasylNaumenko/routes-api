/*
 * telegram: @VasylNaumenko
 */

package osrm

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"routes-api/pkg/api/v1/models"
)

func TestSorting(t *testing.T) {
	have := models.Response{
		Source: "13.388798,52.517033",
		Routes: models.Routes{
			{
				Destination: "13.388798,52.517033",
				Duration:    251.5,
				Distance:    1884.8,
			},
			{
				Destination: "13.397631,52.529430",
				Duration:    372.2,
				Distance:    2946.1,
			},
			{
				Destination: "13.428554,52.523239",
				Duration:    117.6,
				Distance:    950.3,
			},
			{
				Destination: "13.428554,52.523239",
				Duration:    117.6,
				Distance:    800.0,
			},
		},
	}

	want := models.Response{
		Source: "13.388798,52.517033",
		Routes: models.Routes{
			{
				Destination: "13.428554,52.523239",
				Duration:    117.6,
				Distance:    800.0,
			},
			{
				Destination: "13.428554,52.523239",
				Duration:    117.6,
				Distance:    950.3,
			},
			{
				Destination: "13.388798,52.517033",
				Duration:    251.5,
				Distance:    1884.8,
			},
			{
				Destination: "13.397631,52.529430",
				Duration:    372.2,
				Distance:    2946.1,
			},
		},
	}

	sort.Sort(have.Routes)
	assert.Equal(t, want, have)
}
