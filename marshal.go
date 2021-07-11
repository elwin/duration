package duration

import (
	"encoding/json"
)

func (d *Duration) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var durationStr string
	if err := unmarshal(&durationStr); err != nil {
		return err
	}

	durationParsed, err := ParseDuration(durationStr)
	if err != nil {
		return err
	}

	*d = durationParsed

	return nil
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	var durationStr string

	if err := json.Unmarshal(data, &durationStr); err != nil {
		return err
	}

	durationParsed, err := ParseDuration(durationStr)
	if err != nil {
		return err
	}

	*d = durationParsed

	return nil
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}