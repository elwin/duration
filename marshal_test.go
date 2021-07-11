package duration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

var yamlFile = `
minutes: 60m
hours: 24h
days: 30d
weeks: 52w
`

var jsonFile = `
{
	"minutes": "60m",
	"hours": "24h",
	"days": "30d",
	"weeks": "52w"
}
`

type output struct {
	Minutes Duration
	Hours   Duration
	Days    Duration
	Weeks   Duration
}

var unmarshalled = output{
	Minutes: 60 * Minute,
	Hours:   24 * Hour,
	Days:    30 * Day,
	Weeks:   52 * Week,
}

func TestDuration_UnmarshalYAML(t *testing.T) {
	var out output
	assert.NoError(t, yaml.
		NewDecoder(strings.NewReader(yamlFile)).
		Decode(&out),
	)

	assert.Equal(t, unmarshalled, out)
}

func TestDuration_UnmarshalJSON(t *testing.T) {
	var out output
	assert.NoError(t, json.
		NewDecoder(strings.NewReader(jsonFile)).
		Decode(&out),
	)

	assert.Equal(t, unmarshalled, out)
}

func TestDuration_MarshalJSON(t *testing.T) {
	var out output
	assert.NoError(t, json.
		NewDecoder(strings.NewReader(jsonFile)).
		Decode(&out),
	)

	var buf bytes.Buffer
	assert.NoError(t, json.NewEncoder(&buf).Encode(out))
	fmt.Println(buf.String())
}
