/*
 * telegram: @VasylNaumenko
 */

package helpers

import (
	"encoding/json"
	"net/http"
	"time"
)

const defaultTimeout = 10 * time.Second

// GetUrlResponse helps getting data via sending http request
func GetUrlResponse(url string, target interface{}) error {
	httpClient := http.Client{Timeout: defaultTimeout}
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
