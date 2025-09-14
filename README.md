# rd: Unofficial [raindrop.io](https://raindrop.io) command-line interface

rd is a simple command-line interface for [raindrop.io](https://raindrop.io). rd is designed with unix philosophy in mind and it's meant to be used with other tools and scripts.

> **NOTE:** I don't have any plans for adding collection support, since I rarely use them, although backward compatible PR's are always welcome.

```
Usage: rd <command> [<args>]

Options:
  --help, -h             display this help and exit

Commands:
  list                   list raindrops
  add                    add new raindrop
  edit                   edit raindrop
  tags                   get list of tags
  rename-tag             rename tag
  backup                 backup actions
```

## Installation

```console
go install -v github.com/rythoris/rd/cmd/rd@latest
```

Before running the program make sure to create an API token and add it to your [configuration file](#configuration).

## Quick Start

In order to access the [raindrop.io](https://raindrop.io) API you must create an API token. You can create this token in [App Management Console](https://app.raindrop.io/settings/integrations). Since were not using the complicated oauth system, the test token should be enough for our operations. You can consult the [raindrop.io Official Documentation](https://developer.raindrop.io/v1/authentication/token) for more information.

After obtaining the API token you can pass the API token using the `RD_RAINDROPIO_TOKEN` environment variables.

```console
$ RD_RAINDROPIO_TOKEN="YOUR TOKEN HERE" rd list
```

It's not recommended to save the API token in your shell's `rc` file for security reasons. Instead, consider using a password manager to securely store and manage your API tokens. Additionally, if you need to run a command that includes sensitive information, you can prevent it from being saved in the shell history by placing a space before the command. This way, the command won't be recorded in your history. This feature can be enabled in bash by setting the `HISTCONTROL` environment variable to `ignorespace`, consider reading the `bash(1)` man page for more information.

## License

This project is licensed under [BSD-3-Clause](https://opensource.org/license/bsd-3-clause/). See [LICENSE](./LICENSE) for more details.
