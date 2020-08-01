package cmd

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/waznico/atom-reader/internal/pkg/config"
)

func describe(alias string, id string) {
	var fc config.Config
	fc.LoadConfig()

	var feedURI string
	for _, fi := range fc.Feeds.Feeds {
		if fi.Alias == alias {
			feedURI = fi.URI
		}
	}

	if len(feedURI) == 0 {
		fmt.Printf("No feed to alias %v was found!\n", alias)
		return
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedURI)
	if err != nil {
		fmt.Println("Feed couldn't be reached. Please check your network connection and try again.")
	}
	for _, element := range feed.Items {
		if element.GUID == id {
			fmt.Print(element.Title)
			fmt.Print(" | Veröffentlicht am: " + formatDate(element.Published) + "\n")
			fmt.Println(" ")
			fmt.Println(element.Description)
			return
		}
	}
	fmt.Println(id + " not found!")
}

func formatDate(date string) string {
	ref := "2006-01-02T15:04:05-07:00"
	t, err := time.Parse(ref, date)

	if err != nil {
		return ""
	}

	loc, _ := time.LoadLocation("Europe/Berlin")
	return t.In(loc).Format("02.01.2006 um 15:04 Uhr")
}
