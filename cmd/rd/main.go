package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
)

type Cli struct {
	List      *ListCommand      `arg:"subcommand:list"       help:"list raindrops"`
	Add       *AddCommand       `arg:"subcommand:add"        help:"add new raindrop"`
	Edit      *EditCommand      `arg:"subcommand:edit"       help:"edit raindrop"`
	Tags      *TagsCommand      `arg:"subcommand:tags"       help:"get list of tags"`
	RenameTag *RenameTagCommand `arg:"subcommand:rename-tag" help:"rename tag"`
	Backup    *BackupCommand    `arg:"subcommand:backup"     help:"backup actions"`
}

func main() {
	var (
		cli Cli
	)

	token, ok := os.LookupEnv("RD_RAINDROPIO_TOKEN")
	if !ok {
		fmt.Fprintf(os.Stderr, "%s: the 'RD_RAINDROPIO_TOKEN' environment variable is not set.\n", os.Args[0])
		os.Exit(1)
	}

	p := arg.MustParse(&cli)
	switch {
	case cli.List != nil:
		os.Exit(cli.List.Run(token))
	case cli.Add != nil:
		os.Exit(cli.Add.Run(token))
	case cli.Tags != nil:
		os.Exit(cli.Tags.Run(token))
	case cli.RenameTag != nil:
		os.Exit(cli.RenameTag.Run(token))
	case cli.Edit != nil:
		os.Exit(cli.Edit.Run(token))
	case cli.Backup != nil:
		switch {
		case cli.Backup.Create != nil:
			os.Exit(cli.Backup.Create.Run(token))
		case cli.Backup.Download != nil:
			os.Exit(cli.Backup.Download.Run(token))
		case cli.Backup.List != nil:
			os.Exit(cli.Backup.List.Run(token))
		default:
			p.WriteHelp(os.Stderr)
			os.Exit(1)
		}
	default: // sub-command is required
		p.WriteHelp(os.Stderr)
		os.Exit(1)
	}
}
