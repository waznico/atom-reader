package cmd

import (
	"fmt"
	"os"

	"github.com/waznico/atom-reader/internal/pkg/config"

	"github.com/jedib0t/go-pretty/table"

	"github.com/mmcdole/gofeed"
)

func list() {
	var c config.Config

	c.LoadConfig()

	for fi := range c.Feeds.Feeds {
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(c.Feeds.Feeds[fi].URI)
		if err != nil {
			fmt.Printf("Feed %v couldn't be reached. Please check your network connection and try again.\n", c.Feeds.Feeds[fi].Name)
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle("Feeds of %v", c.Feeds.Feeds[fi].Name)
		t.AppendHeader(table.Row{"ID", "Title"})

		for i := 0; i < 10; i++ {
			t.AppendRows([]table.Row{
				{feed.Items[i].GUID, feed.Items[i].Title},
			})
		}
		t.Render()
	}
}
