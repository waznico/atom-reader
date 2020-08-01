package jsonconverter

import (
	"encoding/json"
	"fmt"
)

// Feeds struct is a collection of multiple feed objects
type Feeds struct {
	Feeds []Feed `json:"feeds"`
}

// Feed struct contains the name and the uri of the feed
type Feed struct {
	Alias string `json:"alias"`
	Name  string `json:"name"`
	URI   string `json:"uri"`
}

// Convert converts the input byte array into a Feeds struct if possible
func Convert(input []byte) (Feeds, error) {
	var feeds Feeds
	err := json.Unmarshal(input, &feeds)

	if err != nil {
		fmt.Println("An error occured while converting byte array to Feeds")
	}

	return feeds, err
}

// SetDefaults creates a Feeds struct with default values
func SetDefaults() Feeds {
	var feeds Feeds
	var defaultFeed Feed

	defaultFeed.Alias = "heise_news"
	defaultFeed.Name = "heise News"
	defaultFeed.URI = "https://www.heise.de/rss/heise-atom.xml"

	feeds.Feeds = make([]Feed, 1)
	feeds.Feeds[0] = defaultFeed
	return feeds
}

// ConvertFeedToBytes prepares a byte array which can be written into a file
func ConvertFeedToBytes(feeds Feeds) ([]byte, error) {
	jsonBytes, err := json.Marshal(feeds)
	return jsonBytes, err
}
