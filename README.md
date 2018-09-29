
<p align="center">
  <a href="https://github.com/AnalogJ/tentacle">
  <img width="300" alt="tentacle_view" src="https://rawgit.com/AnalogJ/tentacle/master/logo.svg">
  </a>
</p>

# Tentacle

[![Circle CI](https://img.shields.io/circleci/project/github/AnalogJ/tentacle.svg?style=flat-square)](https://circleci.com/gh/AnalogJ/tentacle)
[![Coverage Status](https://img.shields.io/codecov/c/github/AnalogJ/tentacle.svg?style=flat-square)](https://codecov.io/gh/AnalogJ/tentacle)
[![GitHub license](https://img.shields.io/github/license/AnalogJ/tentacle.svg?style=flat-square)](https://github.com/AnalogJ/tentacle/blob/master/LICENSE)
[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/analogj/tentacle)
[![Go Report Card](https://goreportcard.com/badge/github.com/AnalogJ/tentacle?style=flat-square)](https://goreportcard.com/report/github.com/AnalogJ/tentacle)
[![GitHub release](http://img.shields.io/github/release/AnalogJ/tentacle.svg?style=flat-square)](https://github.com/AnalogJ/tentacle/releases)
[![Github All Releases](https://img.shields.io/github/downloads/analogj/tentacle/total.svg?style=flat-square)](https://github.com/AnalogJ/drawbridge/tentacle)

Retrieve your secrets from wherever they live. Vault/Cyberark/Thycotic/Keychain/Keyring/etc.

## Introduction
Tentacle provides a way to retrieve credentials from your secret manager in a standardized way.
Tentacle has a CLI but can also be used as a Go library.

Tentacle was designed to be used in automation.

## Features

- Single binary available for macOS and linux (windows comming soon). No external dependencies.
- Simple commandline
- Comprehensive configuration file
- Supports multiple secret managment services
- Allows aliasing multiple secret management systems (stage/prod credential storage)


## Getting Started

1. Download the latest release binary from the Release pages for your OS. (Mac, Windows and Linux available)
2. Rename the downloaded binary to `tentacle`
3. Run chmod +x tentacle
4. Move the renamed binary into your PATH, eg. `/usr/bin/local`
5. Run `tentacle --help` from a terminal to confirm it was installed correctly
6. Add a configuration file to `~/tentacle.yaml`. See [Configuration](#configuration) section

## Usage

```bash
$ tentacle help
 ____  ____  __ _  ____  __    ___  __    ____
(_  _)(  __)(  ( \(_  _)/ _\  / __)(  )  (  __)
  )(   ) _) /    /  )( /    \( (__ / (_/\ ) _)
 (__) (____)\_)__) (__)\_/\_/ \___)\____/(____)
github.com/AnalogJ/tentacle                    darwin.amd64-1.0.0

 NAME:
   tentacle - Summary retrieval made simple

USAGE:
   tentacle-darwin-amd64 [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   Jason Kulatunga <jason@thesparktree.com>

COMMANDS:
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug         Enable Debug mode, with extra logging (default: false)
   --output value  Specify output type. Allowed: 'json', 'yaml', 'table', 'raw' (default: "yaml")
   --help, -h      show help (default: false)
   --version, -v   print the version (default: false)

```

## Actions

### Get


### List


## Output formats


## Configuration



## Testing [![Circle CI](https://img.shields.io/circleci/project/github/AnalogJ/tentacle.svg?style=flat-square)](https://circleci.com/gh/AnalogJ/tentacle)
Tentacle provides an extensive test-suite based on `go test`.
You can run all the integration & unit tests with `go test`

CircleCI is used for continuous integration testing: https://circleci.com/gh/AnalogJ/tentacle


## Contributing
If you'd like to help improve Tentacle, clone the project with Git and install dependencies by running:

```
$ git clone git://github.com/AnalogJ/tentacle
$ dep -ensure
```

Work your magic and then submit a pull request. We love pull requests!

If you find the documentation lacking, help us out and update this README.md.
If you don't have the time to work on Drawbridge, but found something we should know about, please submit an issue.





## Current Providers
Tentacle supports retrieving secrets via API's as well as OS specific credential storage.

### Network/Cloud

1. Cyberark
2. Thycotic
3. Conjur


### OS Secret Storage

- MacOS - keychain


## Future Providers

### Network/Cloud
- Dashlane
- Keeper
- RoboForm
- Google SmartLock
- Lastpass
- 1Password
- AWS
- SplashID Safe
- Norton Identity Safe
- Trend Micro Password Manager
- Bitwarden
- PasswordStore (Pass)
- Vault
- Chef

### OS
- Linux - Keyring
- Linux - Secretservice
- Windows - Password vault
- Mac - KeyChain

### Local
- Chrome
- Firefox
- keepass
- gopass
- Password Safe
- Buttercup


## Versioning
We use SemVer for versioning. For the versions available, see the tags on this repository.

## Authors
Jason Kulatunga - Initial Development - @AnalogJ

## License

- Dual-Licensed
    - GPL
    - MIT - Contact jason@thesparktree.com for more information.
- [Logo: Octopus by Alice Noir  from the Noun Project](https://thenounproject.com/search/?q=tentacle&i=593081)


## References

- https://github.com/keybase/go-keychain
- https://github.com/99designs/keyring
- https://github.com/tmc/keyring
- https://github.com/sethvargo/vault-token-helper-osx-keychain
- https://github.com/cloudflare/gokey
- https://github.com/havoc-io/go-keytar
- https://github.com/olekukonko/tablewriter
- https://github.com/lastpass/lastpass-cli
- https://github.com/cyberark/conjur-api-go
- https://github.com/hoop33/go-cyberark
- https://cla-assistant.io/
- https://github.com/cyberark/summon
- https://medium.com/learning-the-go-programming-language/writing-modular-go-programs-with-plugins-ec46381ee1a9
- https://thycotic.force.com/support/s/article/Accessing-Secret-Server-programmatically-Curl
- https://thycotic.com/wp-content/uploads/2014/04/SS_WebServicesGuide.pdf

## Possible Names
- Seer
- invoke
- incant
- despell
- arcane
- shroud
- keyhole - https://thenounproject.com/search/?q=keyhole&i=200437
