package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/waznico/atom-reader/internal/pkg/config"

	"github.com/jedib0t/go-pretty/table"

	"github.com/mmcdole/gofeed"
)

func listAll() {
	var c config.Config
	c.LoadConfig()

	for _, fi := range c.Feeds.Feeds {
		fp := gofeed.NewParser()
		feed, err := fp.ParseURL(fi.URI)
		if err != nil {
			fmt.Printf("Feed %v couldn't be reached. Please check your network connection and try again.\n", fi.Name)
		}
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.SetTitle("Feeds of %v", fi.Name)
		t.AppendHeader(table.Row{"ID", "Title"})

		for i := 0; i < 10; i++ {
			t.AppendRows([]table.Row{
				{feed.Items[i].GUID, feed.Items[i].Title},
			})
		}
		t.Render()
	}
}

func list(alias, amount string) {
	amt, err := strconv.Atoi(amount)

	if err != nil {
		fmt.Printf("The given amount of %v could not be parsed!\n", amount)
		return
	}

	var c config.Config
	c.LoadConfig()

	for _, fi := range c.Feeds.Feeds {
		if fi.Alias == alias {
			fp := gofeed.NewParser()
			feed, parseErr := fp.ParseURL(fi.URI)

			if parseErr != nil {
				fmt.Printf("Feed %v couldn't be reached. Please check your network connection and try again.\n", fi.Name)
			}
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.SetTitle("Feeds of %v", fi.Name)
			t.AppendHeader(table.Row{"ID", "Title"})

			n := 0
			if amt < len(feed.Items) {
				n = amt
			} else {
				n = len(feed.Items)
			}

			for i := 0; i < n; i++ {
				t.AppendRows([]table.Row{
					{feed.Items[i].GUID, feed.Items[i].Title},
				})
			}
			t.Render()
			return
		}
	}
	fmt.Printf("The feed to alias %v could not be found!\n", alias)
}
