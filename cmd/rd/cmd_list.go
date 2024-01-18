package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/rythoris/rd"
	"github.com/rythoris/rd/internal/config"
)

type ListCommand struct {
	Format   PrintFormat `arg:"-f,--format" help:"print fromat (possible values: jsonl, json_array, tsv)" default:"tsv"`
	NoHeader bool        `arg:"-n,--no-header" help:"do not print tsv header (only effects the tsv print format)" default:"false"`
}

func (c ListCommand) Run(config config.Config) int {
	rds, err := rd.GetRaindrops(config.Token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
		return 1
	}

	switch c.Format {
	case FormatTsv:
		if !c.NoHeader {
			fmt.Printf("id\tcreated\tmodified\ttitle\tlink\ttags\n")
		}
		for _, r := range rds {
			fmt.Printf(
				"%d\t%s\t%s\t%s\t%s\t%s\n",
				r.ID,
				r.Created.Format(time.ANSIC),
				r.Modified.Format(time.ANSIC),
				r.Title,
				r.Link,
				strings.Join(r.Tags, ","),
			)
		}
	case FormatJsonLine:
		sb := strings.Builder{}
		for _, r := range rds {
			b, err := json.Marshal(r)
			if err != nil {
				fmt.Fprintf(os.Stderr, "[!] ERROR: marshal error: %s\n", err.Error())
				return 1
			}
			sb.WriteString(string(b) + "\n")
		}
		fmt.Print(sb.String())
	case FormatJsonArray:
		b, err := json.MarshalIndent(rds, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "[!] ERROR: marshal error: %s\n", err.Error())
			return 1
		}
		fmt.Println(string(b))
	}
	return 0
}

type PrintFormat uint

const (
	FormatJsonLine PrintFormat = iota
	FormatJsonArray
	FormatTsv
)

var printFormatNames = [...]string{
	FormatJsonLine:  "jsonl",
	FormatJsonArray: "json_array",
	FormatTsv:       "tsv",
}

func (f *PrintFormat) UnmarshalText(data []byte) error {
	i := slices.Index(printFormatNames[:], strings.ToLower(string(data)))
	if i == -1 {
		return fmt.Errorf("invalid print format: %s", string(data))
	}
	*f = PrintFormat(i)
	return nil
}
