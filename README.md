**cosmocope** (**cosmo**s + teles**cope**) is a tool that helps you discover projects, tools, SDKs, 
libraries and modules that are part of the [Cosmos](https://cosmos.network) Ecosystem

Currently, the only source crawled by the tool is [Github](https://github.com). 
It leverages [Github's search API](https://developer.github.com/v3/search) to discover projects 
tagged with `cosmos-sdk` topic and only public repositories can be crawled. 

In the future, the idea is to engage the community to add Cosmos topics (tags) to their Github public repositories in order to allow 
the tool to discover these repositories as part of the Cosmos Ecosystem.

**NOTE:** Because this tool uses unauthenticated API calls to Github, there's a [rate limit](https://developer.github.com/v3/search/#rate-limit) of __10 requests__ per minutes so running this tool many times per minute might give you error messages.

## Installation

To download and install this tool, please ensure
[Go v1.16 or later is installed](https://golang.org/dl/), then run the following command from a terminal shell:

```shell
$ go get github.com/cosmocope/cosmocope
$ go install github.com/cosmocope/cosmocope
```

## Discover commands

### Discovering projects

This command will use the Github search API to fetch repositories that have the `cosmos-sdk` topic. 

**NOTE:** The list of projects returned is not a "curated" or "official" list of Cosmos projects. Anyone can tag their repository with the `cosmos-sdk` topic. The intention is to use this tool for "discoverability" only. For an official list of Cosmos SDK projects please check the official [Cosmos Network Ecosystem](https://cosmos.network/ecosystem) website.
```shell
$ cosmocope discover projects --help

This command searches for projects on Github that are tagged with the 'cosmos-sdk' topic

Usage:
  cosmocope discover projects [flags]

Flags:
  -h, --help   help for projects
  -j, --json   Output results to JSON

```
Run this command from a terminal shell:

```shell
$ cosmocope discover projects
```


#### JSON output

If you need a JSON output instead of a printed table, use the `--json` or `-j` flag

```shell
$ cosmocope discover projects --json
```

### Discovering modules

This command will fetch all the projects as the command above and will search for a folder name `x` in the repository. This is an indication that this project is based on the `cosmos-sdk` and the `x` folder contains modules. 

This command usually takes under 30 seconds to return the results.

**NOTE:** This command doesn't guarantee that the folders shown are actually Cosmos SDK modules. The intention is to use this tool for "discoverability" only. For an official list of modules, please check [Atlas](https://atlas.cosmos.network) which is the official Cosmos SDK module registry.

```shell
$ cosmocope discover projects --help

This command searches for projects on Github that are tagged with the 'cosmos-sdk' topic

Usage:
  cosmocope discover projects [flags]

Flags:
  -h, --help   help for projects
  -j, --json   Output results to JSON

```
Run this command from a terminal shell:

```shell
$ cosmocope discover modules
```

#### JSON output

If you need a JSON output instead of a printed table, use the `--json` or `-j` flag

```shell
$ cosmocope discover modules --json
```

### FAQ


### Disclaimer