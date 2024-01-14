package main

import (
	"fmt"
	"os"

	"github.com/rythoris/rd"
	"github.com/rythoris/rd/internal/config"
)

type EditCommand struct {
	ID      int      `arg:"positional,required" help:"link id"`
	NewLink string   `arg:"-l,--link" help:"new link"`
	Tags    []string `arg:"-t,--tag,separate" help:"tags"`
}

func (c EditCommand) Run(config config.Config) int {
	// if the tags and newlink is empty just delete it.
	if c.NewLink == "" && len(c.Tags) == 0 {
		if err := rd.RemoveRaindrop(config.Token, c.ID); err != nil {
			fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
			return 1
		}
	} else {
		if err := rd.UpdateRaindrop(config.Token, c.ID, c.NewLink, c.Tags); err != nil {
			fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
			return 1
		}
	}
	return 0
}
