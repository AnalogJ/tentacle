
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

1. [Cyberark](https://www.cyberark.com/)
1. [Conjur](https://www.conjur.org/)
1. [Lastpass](https://www.lastpass.com)
1. [Thycotic](https://thycotic.com/)

### OS Secret Storage

- MacOS - keychain


## Future Providers

### Network/Cloud
- [1Password](https://1password.com/)
- [Avecto Defendpoint](https://www.avecto.com/defendpoint)
- AWS
- [BeyondTrust PowerBroker](https://www.beyondtrust.com/products/powerbroker-password-safe/)
- [Bitwarden](https://bitwarden.com/)
- [BomGar PAM](https://www.bomgar.com/)
- [Chef Encrypted Databags](https://docs.chef.io/data_bags.html)
- [Dashlane](https://www.dashlane.com/)
- [Devolutions Password Server](https://devolutions.net/products/password-server)
- [Google SmartLock](https://developers.google.com/identity/smartlock-passwords/case-studies)
- [IT Glue Password Vault](https://itglue.com/features/password-vault/)
- [Keeper](https://keepersecurity.com/)
- [KronTech Single Connect](https://krontech.com/)
- [Manage Engine Password Manager Pro](https://www.manageengine.com/products/passwordmanagerpro/)
- [Norton Identity Safe](https://my.norton.com/extspa/idsafe)
- [Passwordstate](https://www.clickstudios.com.au/)
- [PasswordStore (Pass)](https://www.passwordstore.org/)
- [Pleasant Password Server](http://www.pleasantsolutions.com/passwordserver)
- [RoboForm](https://www.roboform.com/)
- [SplashID Safe](https://splashid.com/)
- [Teampass](https://github.com/nilsteampassnet/TeamPass) - Keepass for Teams
- [Trend Micro Password Manager](https://pwm.trendmicro.com/)
- [Vault](https://www.vaultproject.io/)
- [XTon PAM](https://www.xtontech.com/overview/features/)

### OS
- [Linux - Keyring](http://man7.org/linux/man-pages/man7/keyrings.7.html)
- [Linux - Secretservice](https://community.kde.org/KDE_Utils/ksecretsservice)
- [Windows - Credential Manager/ Password vault](https://support.microsoft.com/en-us/help/4026814/windows-accessing-credential-manager)
- [Mac - KeyChain](https://support.apple.com/guide/keychain-access/what-is-keychain-access-kyca1083/mac)

### Local
- [Chrome](https://support.google.com/chrome/answer/95606?hl=en&co=GENIE.Platform%3DDesktop)
- [Firefox](https://support.mozilla.org/en-US/kb/password-manager-remember-delete-change-and-import)
- [keepass](https://keepass.info/)
- [gopass](https://github.com/gopasspw/gopass)
- [Password Safe](https://pwsafe.org/)
- [Buttercup](https://buttercup.pw/)


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
- https://github.com/robertknight/1pass
- https://discussions.agilebits.com/discussion/89740/api-to-pull-1password-information
- https://github.com/1Password/1password-teams-open-source

## Possible Names
- Seer
- invoke
- incant
- despell
- arcane
- shroud
- keyhole - https://thenounproject.com/search/?q=keyhole&i=200437
