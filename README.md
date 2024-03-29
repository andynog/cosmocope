# cosmocope

[![Go](https://github.com/andynog/cosmocope/actions/workflows/go.yml/badge.svg)](https://github.com/andynog/cosmocope/actions/workflows/go.yml)

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

```
go install github.com/andynog/cosmocope/v2@latest
```

Please ensure you have the `GOBIN` environment variable properly set.

## Help and Usage 

To find more information about the tool help and usage run the following command:

```
cosmocope discover --help
```

A help message will be displayed:

```
$ cosmocope discover --help
Find repositories, releases, and modules for projects in the Cosmos ecosystem.

Usage:
  cosmocope discover [flags]
  cosmocope discover [command]

Available Commands:
  modules     Find Cosmos SDK modules
  projects    Find Cosmos projects
  releases    Find Github repository releases

Flags:
  -h, --help   help for discover

Use "cosmocope discover [command] --help" for more information about a command.
```

## Discover commands

### Discovering projects

This command will use the Github search API to fetch repositories that have the `cosmos-sdk` topic. The results will be sorted by the last updated date (commit), the most recent will show on top.

**NOTE:** The list of projects returned is not a "curated" or "official" list of Cosmos projects. Anyone can tag their repository with the `cosmos-sdk` topic. The intention is to use this tool for "discoverability" only. For an official list of **Cosmos SDK projects** please check the official [Cosmos Network Ecosystem](https://cosmos.network/ecosystem) website.

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
cosmocope discover projects
```

### Discovering modules

This command will fetch all the projects as the command above and will search for a folder name `x` in the repository. This is an indication that this project is based on the `cosmos-sdk` and the `x` folder contains modules. 

This command usually takes under 30 seconds to return the results. This command executes a lot of Github API call in order to retrieve information about the modules folders. You might be able to run it a couple of times each hour before reaching the rate limit. If you reach the rate limit, the command will tell how long you have to wait to try it again.

**NOTE:** This command doesn't guarantee that the folders shown are actually Cosmos SDK modules. The intention is to use this tool for "discoverability" only. For an official list of modules, please check [Atlas](https://atlas.cosmos.network) which is the official **Cosmos SDK module registry**.

```
$ cosmocope discover modules --help

This command searches for projects on Github that are tagged 
with the 'cosmos-sdk' topic and the tool crawls each repository 
looking for a folder named 'x' in the repository.

Usage:
  cosmocope discover modules [flags]

Flags:
  -h, --help   help for modules
  -j, --json   Output results to JSON
```

Run this command from a terminal shell:

```
cosmocope discover modules
```


### Discovering releases

This command will fetch all the releases for a Githube repository. 

```
$ cosmocope discover releases --help

This command lists the releases for a Github repository.

Usage:
  cosmocope discover releases [flags]

Flags:
  -h, --help          help for releases
  -j, --json          Output results to JSON
  -r, --repo string   Specify the Github Repository URL (e.g. https://github.com/cosmos/cosmos-sdk
```

Run this command from a terminal shell:

```
cosmocope discover releases
```

### JSON output

If you need a JSON output instead of a printed table, use the `--json` or `-j` flag. The JSON output also might include additional data (e.g. description and forks) that is not shown in the default output (table)

* Projects
  
    ```
    cosmocope discover projects --json
    ```

* Modules
  
    ```
   cosmocope discover modules --json
   ```
  
* Releases 

    ```
   cosmocope discover releases --json
   ```

### FAQ

For some Frequently Asked Questions please visit the [FAQ](FAQ.md) page.

### Disclaimer

The information provided by this tool is directly retrieved from Github using its Search API. The tool does not have any control on the content returned by Github. The tool only displays the results in a nice format (table format or Json).

Please exercise due diligence in assessing the results returned.
