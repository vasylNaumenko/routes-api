/*
 * telegram: @VasylNaumenko
 */

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckStringParam(t *testing.T) {
	// Empty value case
	{
		value := ""
		checkStringParam(&value, "default")
		assert.Equal(t, "default", value)
	}
	// Non-empty value case
	{
		value := "test"
		checkStringParam(&value, "default")
		assert.NotEqual(t, "default", value)
	}
}
