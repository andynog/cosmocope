# cosmoscope - A Cosmos Telescope tool

This tool allows you to crawl [Github](https://github.com) in order to discover [Cosmos](https://cosmos.network) based projects.

Currently, the only source crawled by the tool is Github. It leverages Github's search API to discover projects tagged with 'cosmos-sdk' topic and only public repositories can be crawled.

## Installing the tool

In order to run the tool, it is assumed that you have a recent version of (Golang)[https://golang.org/doc/install] properly setup on your machine.

```$ go install github.com/andynog/cosmoscope```

## Crawling to find projects

1. Open a terminal
2. Run the tool

    ```$ cosmoscope```
