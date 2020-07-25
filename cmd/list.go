package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"

	"github.com/mmcdole/gofeed"
)

func list() {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://www.heise.de/rss/heise-atom.xml")
	if err != nil {
		fmt.Println("Feed couldn't be reached. Please check your network connection and try again.")
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Title"})

	for i := 0; i < 10; i++ {
		t.AppendRows([]table.Row{
			{feed.Items[i].GUID, feed.Items[i].Title},
		})
	}
	t.Render()
}
