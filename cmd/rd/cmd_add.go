package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/asaskevich/govalidator"
	"github.com/rythoris/rd"
	"github.com/rythoris/rd/internal/config"
)

var (
	// tags regexp
	TagRegexp = regexp.MustCompile(`^[a-zA-Z0-9-_/]+$`)
)

type AddCommand struct {
	Link string   `arg:"positional,required" help:"link"`
	Tags []string `arg:"positional" help:"tags"`
}

func (c AddCommand) Run(config config.Config) int {
	if !govalidator.IsURL(c.Link) {
		fmt.Fprintf(os.Stderr, "[!] ERROR: invalid url: %s\n", c.Link)
		return 1
	}

	for _, t := range c.Tags {
		if !TagRegexp.MatchString(t) {
			fmt.Fprintf(os.Stderr, "[!] ERROR: bad tag: %s\n", t)
			return 1
		}
	}

	id, err := rd.CreateRaindrop(config.Token, c.Link, c.Tags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
		return 1
	}
	fmt.Printf("%d\n", id)
	return 0
}
