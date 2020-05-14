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

### Installation

TBD

### Usage

Add a script below on your shell startup file like `~/.bash_profile`, `~/.zshrc` `~/.profile`.

```sh
eval "$(mirrorctl export --location <location>)"
```
