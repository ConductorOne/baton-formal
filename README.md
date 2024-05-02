![Baton Logo](./docs/images/baton-logo.png)

# `baton-formal` [![Go Reference](https://pkg.go.dev/badge/github.com/conductorone/baton-formal.svg)](https://pkg.go.dev/github.com/conductorone/baton-formal) ![main ci](https://github.com/conductorone/baton-formal/actions/workflows/main.yaml/badge.svg)

`baton-formal` is a connector for [Formal](https://joinformal.com) built using the [Baton SDK](https://github.com/conductorone/baton-sdk). It communicates with the Formal API to sync data about users and groups.

Check out [Baton](https://github.com/conductorone/baton) to learn more about the project in general.

# Prerequisites

To work with the connector, you will require a Formal API key. You can obtain this key by signing up for a Formal account and generating an API key in the Formal dashboard.
# Getting Started

## brew

```
brew install conductorone/baton/baton conductorone/baton/baton-formal

BATON_FORMAL_API_KEY=<api-key> baton-formal
baton resources
```

## docker

```
docker run --rm -v $(pwd):/out -e BATON_FORMAL_API_KEY=<api-key> ghcr.io/conductorone/baton-formal:latest -f "/out/sync.c1z"
docker run --rm -v $(pwd):/out ghcr.io/conductorone/baton:latest -f "/out/sync.c1z" resources
```

## source

```
go install github.com/conductorone/baton/cmd/baton@main
go install github.com/conductorone/baton-formal/cmd/baton-formal@main

BATON_FORMAL_API_KEY=<api-key> baton-formal
baton resources
```

# Data Model

`baton-formal` will pull down information about the following Formal resources:

- Users
- Groups

# Contributing, Support and Issues

We started Baton because we were tired of taking screenshots and manually building spreadsheets. We welcome contributions, and ideas, no matter how small -- our goal is to make identity and permissions sprawl less painful for everyone. If you have questions, problems, or ideas: Please open a Github Issue!

See [CONTRIBUTING.md](https://github.com/ConductorOne/baton/blob/main/CONTRIBUTING.md) for more details.

# `baton-formal` Command Line Usage

```
baton-formal

Usage:
  baton-formal [flags]
  baton-formal [command]

Available Commands:
  capabilities       Get connector capabilities
  completion         Generate the autocompletion script for the specified shell
  help               Help about any command

Flags:
      --client-id string        The client ID used to authenticate with ConductorOne ($BATON_CLIENT_ID)
      --client-secret string    The client secret used to authenticate with ConductorOne ($BATON_CLIENT_SECRET)
  -f, --file string             The path to the c1z file to sync with ($BATON_FILE) (default "sync.c1z")
      --formal-api-key string   Your Formal API KEY ($BATON_FORMAL_API_KEY)
  -h, --help                    help for baton-formal
      --log-format string       The output format for logs: json, console ($BATON_LOG_FORMAT) (default "json")
      --log-level string        The log level: debug, info, warn, error ($BATON_LOG_LEVEL) (default "info")
  -p, --provisioning            This must be set in order for provisioning actions to be enabled. ($BATON_PROVISIONING)
  -v, --version                 version for baton-formal

Use "baton-formal [command] --help" for more information about a command.

```
