package playback

import (
	"encoding/json"
	"fmt"

	"github.com/StageAutoControl/controller/pkg/disk"
)

// EnsureDefaultConfig ensures that the default configuration exists in given storage
func EnsureDefaultConfig(storage storage) error {
	config := &Config{}
	if err := storage.Read(paramsStorageKey, config); err != nil {
		if err != disk.ErrNotExists {
			return fmt.Errorf("failed to find playback config: %v", err)
		}

		if err := json.Unmarshal([]byte(defaultConfig), config); err != nil {
			return fmt.Errorf("failed to decode the default config: %v", err)
		}

		if err := storage.Write(paramsStorageKey, config); err != nil {
			return fmt.Errorf("failed to write the default config to storage: %v", err)
		}
	}

	return nil
}
