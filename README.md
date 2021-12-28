# OpenArch CLI

![OpenArch][openarch-logo]

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/open-arch/oarch/Build?longCache=tru&label=Build&logo=github%20actions&logoColor=fff&style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/open-arch/oarch?style=flat-square)](https://goreportcard.com/report/github.com/open-arch/oarch)

The OpenArch CLI creates, manages, and publishes your OpenArch projects.

- [OpenArch CLI](#openarch-cli)
  - [Installation](#installation)
    - [go install](#go-install)
  - [Initializes a new project](#initializes-a-new-project)
  - [Built With](#built-with)

## Installation

### go install

```bash
go install github.com/open-arch/oarch@latest
```

## Initializes a new project

```bash
mkdir my-project
cd ./my-project/
oarch init
```

## Built With

- [golang][golang] - The language
- [cobra][cobra-cli] - The CLI framework

[cobra-cli]: https://github.com/spf13/cobra
[golang]: https://go.dev/
[openarch-logo]: .attachments/openarch-logo-extended.svg
