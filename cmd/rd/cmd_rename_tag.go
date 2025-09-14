package main

import (
	"fmt"
	"os"

	"github.com/rythoris/rd"
)

type RenameTagCommand struct {
	Tag    string `arg:"positional,required" help:"name of the tag"`
	NewTag string `arg:"positional,required" help:"new name for the tag"`
}

func (c RenameTagCommand) Run(token string) int {
	if err := rd.RenameTag(token, c.Tag, c.NewTag); err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
		return 1
	}
	return 0
}
