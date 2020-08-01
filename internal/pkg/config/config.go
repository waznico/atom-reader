package config

import (
	"fmt"
	"os"
	"path"

	"github.com/waznico/atom-reader/internal/pkg/json"
	"github.com/waznico/atom-reader/internal/pkg/jsonconverter"
)

// Config holds all settings
type Config struct {
	Feeds jsonconverter.Feeds
}

// LoadConfig loads all config files into the Config struct
func (c *Config) LoadConfig() {
	wdir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while identifying current working directory. Config could not be loaded.")
		return
	}

	feedsJSONPath := path.Join(wdir, "feeds.json")
	if _, err := os.Stat(feedsJSONPath); err != nil {
		createFeedsJSON(feedsJSONPath)
	}

	byteContent, err := json.ReadJSON(feedsJSONPath)
	if err == nil {
		feeds, convErr := jsonconverter.Convert(byteContent)

		if convErr == nil {
			c.Feeds = feeds
		}
	}
}

func createFeedsJSON(filePath string) error {
	defaults := jsonconverter.SetDefaults()
	buffer, convErr := jsonconverter.ConvertFeedToBytes(defaults)
	if convErr != nil {
		return convErr
	}

	writeErr := json.WriteJSON(filePath, buffer)

	if writeErr != nil {
		return writeErr
	}
	return nil
}
