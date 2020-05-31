# cosmocope

_cosmocope_ (Cosmos + Telescope) is a tool that helps you discover projects, tools, SDKs, libraries and modules that are part of the Cosmos Ecosystem

Currently, the only source crawled by the tool is [Github](https://github.com). 
It leverages Github's search API to discover projects tagged with 'cosmos-sdk' topic 
and only public repositories can be crawled. In the future, the idea is to engage the 
community to add Cosmos topics to their Github public repositories in order to allow 
the tool to discover their repositories. 

## Installation

To download and install this tool, please ensure
[Go v1.14 or later is installed](https://golang.org/dl/), then run the following command from a terminal shell:

```shell
$ go get github.com/andynog/cosmocope
```

## Discovering projects

Run this command from a terminal shell:

    ```$ cosmocope discover projects```

### JSON output

If you need a JSON output instead of a printed table, use the `--json` or `-j` flag

    ```$ cosmocope discover projects --json```