package models

import (
	"encoding/json"
	"fmt"
)

func (c *UserCache) unmarshallCacheData(cache map[string]string) error {
	jsonData, err := json.Marshal(cache)
	if err != nil {
		return fmt.Errorf("failed to marshal cache to JSON: %v", err)
	}

	if unmarshalErr := json.Unmarshal(jsonData, c); unmarshalErr == nil {
		return fmt.Errorf("failed to unmarshal cache JSON: %v", unmarshalErr)
	}

	return nil
}
