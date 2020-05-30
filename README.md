# cosmoscope - A Cosmos Telescope tool

This tool allows you to crawl [Github](https://github.com) in order to discover [Cosmos](https://cosmos.network) based projects.

Currently, the only source crawled by the tool is Github. It leverages Github's search API to discover projects tagged with 'cosmos-sdk' topic and only public repositories can be crawled.

## Installation

To download and install this tool, please ensure
[Go v1.14 or later is installed](https://golang.org/dl/), then run the please run the following command from a terminal shell:

```shell
$ go get github.com/andynog/cosmoscope
```

## Crawling to find projects

1. Open a terminal
2. Run the tool

    ```$ cosmoscope```
