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

## Configuration

rd by default uses the [adrg/xdg](https://github.com/adrg/xdg) library for handling the configuration path.

Here is an example configuration (`$XDG_CONFIG_HOME/rd/config.json`):
```json
{
  "api-token": "your api token"
}
```

The `api-token` is your test token. You can create this token in [App Management Console](https://app.raindrop.io/settings/integrations).

## License

This project is licensed under [BSD-3-Clause](https://opensource.org/license/bsd-3-clause/). See [LICENSE](./LICENSE) for more details.
