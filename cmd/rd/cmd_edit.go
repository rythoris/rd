package main

import (
	"fmt"
	"os"

	"github.com/rythoris/rd"
)

type EditCommand struct {
	ID      int      `arg:"positional,required" help:"link id"`
	NewLink string   `arg:"-l,--link" help:"new link"`
	Tags    []string `arg:"-t,--tag,separate" help:"tags"`
}

func (c EditCommand) Run(token string) int {
	// if the tags and newlink is empty just delete it.
	if c.NewLink == "" && len(c.Tags) == 0 {
		if err := rd.RemoveRaindrop(token, c.ID); err != nil {
			fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
			return 1
		}
	} else {
		if err := rd.UpdateRaindrop(token, c.ID, c.NewLink, c.Tags); err != nil {
			fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
			return 1
		}
	}
	return 0
}
