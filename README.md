# mirrorctl

Manage mirror urls for many repositories.

## Supporting Repositories

- Bundler (Ruby)
- Pipenv (Python)
- Pypi (Python)

## Supporting Locations

If you set an invalid location code, it will be set to default repository urls.

- Korea (`kr`, `korea`)

## Getting Started

### Install

#### Using Homebrew

```sh
brew tap hellodhlyn/mirrorctl
brew install mirrorctl
```

#### Manual Installation

Download a binary from the [releases page](https://github.com/hellodhlyn/mirrorctl/releases).

```sh
tar -xzf mirrorctl_XXX.tar.gz
cp mirrorctl /usr/bin
```

### Usage

Add a script below on your shell startup file like `~/.bash_profile`, `~/.zshrc`, or `~/.profile`.

```sh
# Set your own location code instead of <location>.
eval "$(mirrorctl export --location <location>)"
```

## Development

### Prerequisites

- Go 1.14 (or greater)
- [Pkger](https://github.com/markbates/pkger)

### Build

```sh
# Build binary
pkger -include /mirrorlist
go build -o dist/mirrorctl

# Run command
./dist/mirrorctl --help
```
