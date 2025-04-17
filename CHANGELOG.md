# Changelog

## v0.1.0

Updated Go version to 1.24 and refreshed dependencies. Improved linting configuration and GitHub Actions workflows, including Dependabot integration.

### BREAKING CHANGES

- bumped Go version to 1.24

### IMPROVEMENTS

- adding dependabot workflow for go dependencies (#12)
- update Go build GitHub action (#13)
- update golangci-lint-action version (#14)
- adding lint config and fixing linting issues (#16)

## v0.0.3

### FEATURES

- Implemented `discover releases` command ([#1])

### IMPROVEMENTS

- Detect Cosmos chains based on go dependency ([#2])
- Show number of forks in discover projects table view ([#4])
- Allow sorting projects by number of forks or stars ([#5])

### BUG FIXES

[#1]: https://github.com/andynog/cosmocope/issues/1
[#2]: https://github.com/andynog/cosmocope/issues/2
[#4]: https://github.com/andynog/cosmocope/issues/4
[#5]: https://github.com/andynog/cosmocope/issues/5

## v0.0.2
*March 16, 2021*

- Fixed dependencies issues

## v0.0.1
*March 7, 2021*

This is the first `cosmocope` official release. This initial release allows users to discover projects and modules in the __Cosmos Ecosystem__ using the __Github__ API.
