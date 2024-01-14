package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"

	"github.com/rythoris/rd/internal/config"
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
		c   config.Config
		err error
	)

	c, err = config.ParseConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] ERROR: %s\n", err.Error())
		os.Exit(1)
	}

	p := arg.MustParse(&cli)
	switch {
	case cli.List != nil:
		os.Exit(cli.List.Run(c))
	case cli.Add != nil:
		os.Exit(cli.Add.Run(c))
	case cli.Tags != nil:
		os.Exit(cli.Tags.Run(c))
	case cli.RenameTag != nil:
		os.Exit(cli.RenameTag.Run(c))
	case cli.Edit != nil:
		os.Exit(cli.Edit.Run(c))
	case cli.Backup != nil:
		switch {
		case cli.Backup.Create != nil:
			os.Exit(cli.Backup.Create.Run(c))
		case cli.Backup.Download != nil:
			os.Exit(cli.Backup.Download.Run(c))
		case cli.Backup.List != nil:
			os.Exit(cli.Backup.List.Run(c))
		default:
			p.WriteHelp(os.Stderr)
			os.Exit(1)
		}
	default: // sub-command is required
		p.WriteHelp(os.Stderr)
		os.Exit(1)
	}
}
