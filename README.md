# git-credential-1password

[![license](https://img.shields.io/github/license/develerik/git-credential-1password.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/develerik/git-credential-1password)](https://goreportcard.com/report/github.com/develerik/git-credential-1password)
[![CodeQL](https://github.com/develerik/git-credential-1password/workflows/CodeQL/badge.svg)](https://github.com/develerik/git-credential-1password/actions?query=workflow%3ACodeQL)
[![Code Style](https://github.com/develerik/git-credential-1password/workflows/Lint%20Code%20Base/badge.svg)](https://github.com/develerik/git-credential-1password/actions?query=workflow%3A%22Lint+Code+Base%22)

Helper to store git credentials inside 1password.

## Table of Contents

- [Installation](#installation)
  - [Dependencies](#dependencies)
  - [Arch Linux](#arch-linux)
  - [From Source](#from-source)
- [Usage](#usage)
- [Support](#support)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [Maintainers](#maintainers)
- [Acknowledgements](#acknowledgements)
- [License](#license)

## Installation

### Dependencies

To use this helper you need to install the 1password cli tool (>=1.8.0) ([download](https://support.1password.com/command-line-getting-started/#set-up-the-command-line-tool))
and of course git.  
You also need to setup the cli tool with your 1password account ([guide](https://support.1password.com/command-line-getting-started/#get-started-with-the-command-line-tool)).

### Arch Linux

On Arch Linux the following packages are available at the AUR:

- `git-credential-1password`: The latest release
- `git-credential-1password-bin`: The latest release (prebuild)
- `git-credential-1password-git`: Builds the current `main` branch
- `git-credential-1password-dev-git`: Builds the current `develop` branch

### From Source

```shell script
git clone https://github.com/develerik/git-credential-1password.git
cd git-credential-1password
make credential-helper
```

Move the built binary (inside the `bin` directory) to somewhere in your PATH.

## Usage

```shell script
git config --global credential.helper '!git-credential-1password'
```

## Support

This project is maintained by [@develerik](https://github.com/develerik). Please understand that we won't be able to
provide individual support via email. We also believe that help is much more valuable if it's shared publicly, so that
more people can benefit from it.

- [**Report a bug**](https://github.com/develerik/git-credential-1password/issues/new?labels=bug&template=bug_report.md)
- [**Requests a new feature**](https://github.com/develerik/git-credential-1password/issues/new?labels=enhancement&template=feature_request.md)
- [**Report a security vulnerability**](https://github.com/develerik/git-credential-1password/issues/new?labels=vulnerability&template=vulnerability_report.md)

## Roadmap

- Support selecting the vault to use (maybe even multiple)
- Maybe an interactive mode for each operation
<!--No changes are currently planned.-->

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct.

## Maintainers

- **Erik Bender** - *Initial work* - [develerik](https://github.com/develerik)

See also the list of [contributors](https://github.com/develerik/git-credential-1password/graphs/contributors) who participated in this project.

## Acknowledgements

- [1Password](https://1password.com) for their awesome [cli tool](https://1password.com/downloads/command-line)
- [Steve (acahir)](https://github.com/acahir) for his [python implementation](https://github.com/acahir/git-credential-1password)
of a 1password credential helper which inspired me to create this project
- [Netlify](https://www.netlify.com) for their [netlify credential helper](https://github.com/netlify/netlify-credential-helper)
implemented in Go which helped me a lot on my own implementation

## License

Distributed under the ISC License. See [LICENSE](LICENSE) for more information.
