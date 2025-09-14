package main

import (
	"fmt"
	"os"

	"github.com/rythoris/rd"
)

type TagsCommand struct {
	Count bool `arg:"-c,--count" help:"print count of each tag"`
}

func (c TagsCommand) Run(token string) int {
	tags, err := rd.GetTags(token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
		return 1
	}

	for _, t := range tags {
		if c.Count {
			fmt.Printf("%d\t", t.Count)
		}
		fmt.Printf("%s", t.Name)
		fmt.Printf("\n")
	}

	return 0
}
