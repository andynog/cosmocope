**cosmocope** (**cosmo**s + teles**cope**) is a tool that helps you discover projects, tools, libraries and modules that are part of the [Cosmos](https://cosmos.network) Ecosystem

Currently, the only source crawled by the tool is [Github](https://github.com). 
It leverages [Github's search API](https://developer.github.com/v3/search) to discover projects 
tagged with `cosmos-sdk` topic and only public repositories can be crawled. 

In the future, the idea is to engage the community to add Cosmos topics (tags) to their Github public repositories in order to allow 
the tool to discover these repositories as part of the Cosmos Ecosystem.

**NOTE:** Because this tool uses unauthenticated API calls to Github, there's a [rate limit](https://developer.github.com/v3/search/#rate-limit) so running this tool many times per minute might give you error messages. The rate limit error message gives you a wait time (in minutes) that you have to wait to run it again successfully.

## Installation

To download and install this tool, please ensure
[Go v1.16 or later is installed](https://golang.org/dl/), then run the following command from a terminal shell:

```shell
$ go get github.com/cosmocope/cosmocope
$ go install github.com/cosmocope/cosmocope
```

## Discover commands

### Discovering projects

This command will use the Github search API to fetch repositories that have the `cosmos-sdk` topic. The results will be sorted by the last updated date (commit), the most recent will show on top.

**NOTE:** The list of projects returned is not a "curated" or "official" list of Cosmos projects. Anyone can tag their repository with the `cosmos-sdk` topic. The intention is to use this tool for "discoverability" only. For an official list of Cosmos SDK projects please check the official [Cosmos Network Ecosystem](https://cosmos.network/ecosystem) website.

```
$ cosmocope discover projects --help

This command searches for projects on Github that are tagged with the 'cosmos-sdk' topic

Usage:
  cosmocope discover projects [flags]

Flags:
  -h, --help   help for projects
  -j, --json   Output results to JSON

```
Run this command from a terminal shell:

```
$ cosmocope discover projects
```


#### JSON output

If you need a JSON output instead of a printed table, use the `--json` or `-j` flag

```
$ cosmocope discover projects --json
```

### Discovering modules

This command will fetch all the projects as the command above and will search for a folder name `x` in the repository. This is an indication that this project is based on the `cosmos-sdk` and the `x` folder contains modules. 

This command usually takes under 30 seconds to return the results.

**NOTE:** This command doesn't guarantee that the folders shown are actually Cosmos SDK modules. The intention is to use this tool for "discoverability" only. For an official list of modules, please check [Atlas](https://atlas.cosmos.network) which is the official Cosmos SDK module registry.

```
$ cosmocope discover projects --help

This command searches for projects on Github that are tagged with the 'cosmos-sdk' topic

Usage:
  cosmocope discover projects [flags]

Flags:
  -h, --help   help for projects
  -j, --json   Output results to JSON

```

Run this command from a terminal shell:

```
$ cosmocope discover modules
```

#### JSON output

If you need a JSON output instead of a printed table, use the `--json` or `-j` flag

```
$ cosmocope discover modules --json
```

### FAQ

For some Frequently Asked Questions please visit the [FAQ](FAQ.md) page.

### Disclaimer

The information provided by this tool is directly retrieved from Github using its Search API. The tool does not have any control on the content returned by Github. The tool only displays the results in a nice format (table format or Json).

Please exercise due diligence in assessing the results returned.