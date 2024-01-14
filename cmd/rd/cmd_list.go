package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	"github.com/rythoris/rd"
	"github.com/rythoris/rd/internal/config"
)

type ListCommand struct {
	Raw bool `arg:"-r,--raw" help:"print raw data as tsv"`
}

func (c ListCommand) Run(config config.Config) int {
	rds, err := rd.GetRaindrops(config.Token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
		return 1
	}

	if c.Raw {
		for _, r := range rds {
			fmt.Printf("%d\t%s\t%s\t%s\n", r.ID, r.Title, r.Link, strings.Join(r.Tags, ","))
		}
	} else {
		t := table.New().
			Border(lipgloss.HiddenBorder()).
			Headers("ID", "Title", "URL", "Tags")
		for _, r := range rds {
			t.Row(fmt.Sprintf("%d", r.ID), r.Title, r.Link, strings.Join(r.Tags, ","))
		}
		fmt.Println(t)
	}
	return 0
}
