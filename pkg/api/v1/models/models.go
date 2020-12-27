/*
 * telegram: @VasylNaumenko
 */

package models

type Route struct {
	Destination string  `json:"destination"`
	Duration    float32 `json:"duration"`
	Distance    float32 `json:"distance"`
}

// Routes implements sort.Interface based on the Duration and Distance field if Duration is equal
type Routes []Route

func (a Routes) Len() int { return len(a) }
func (a Routes) Less(i, j int) bool {
	if a[i].Duration == a[j].Duration {
		return a[i].Distance < a[j].Distance
	}
	return a[i].Duration < a[j].Duration
}
func (a Routes) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// used for osrm response
type Response struct {
	Source string `json:"source"`
	Routes Routes `json:"routes"`
}

// OSRM RESPONSE
// used for osrm request answer
type OsrmResponse struct {
	Code      string       `json:"code"`
	Message   string       `json:"message"`
	Waypoints []Waypoint   `json:"waypoints"`
	Routes    []OsrmRoutes `json:"routes"`
}

type Waypoint struct {
	Distance float32   `json:"distance"`
	Location []float32 `json:"location"`
}

type OsrmRoutes struct {
	Legs []Leg `json:"legs"`
}

type Leg struct {
	Distance float32 `json:"distance"`
	Duration float32 `json:"duration"`
}
