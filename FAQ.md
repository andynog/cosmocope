# Frequently Asked Questions


- [Why a Cosmos project is not showing on the results?](#why-a-cosmos-project-is-not-showing-on-the-results)

- [Are these projects displayed official Cosmos SDK projects?](#are-these-projects-displayed-official-cosmos-sdk-projects)

- [Why use only one topic to discover projects?](#why-use-only-one-topic-to-discover-projects)

- [Why a rate limit error is returned when running a command?](#why-a-rate-limit-error-is-returned-when-running-a-command)


## Why a Cosmos project is not showing on the results?

Currently, the `cosmocope` tool only retrieves Github repositories with a topic named `cosmos-sdk` associated with it. Please add the topic and rerun the tool. It should show up the next time the command `cosmocope discover projects` is executed.

## Are these projects displayed official Cosmos SDK projects?

The projects returned are not part of a curated list of Cosmos projects. If someone creates a project that is not based in the `cosmos-sdk` and decides to tag them with a `cosmos-sdk` project, there is nothing to prevent that since this is a Github feature available to anyone who has access to Github.

Please note that the repositories returned might be something other than a Cosmos SDK chain. Many developers use the `cosmic-sdk` topic to tag their projects as part of the Cosmos ecosystem, so the list returned might include projects for wallets or client libraries, for example, and not only Cosmos chains.

In the future, there might be better ways to list the real Cosmos SDK repositories, possibly based on a vetted or curated list of repositories.

This tool's intention is only to allow the discoverability of new projects that might become part of the Cosmos Ecosystem. For example, the tool might be useful for people who manage or host a list of Cosmos projects.

For an official list of Cosmos SDK projects, please check the official [Cosmos Network Ecosystem](https://cosmos.network/ecosystem) website.

For an official list of modules, please check [Atlas](https://atlas.cosmos.network), the official Cosmos SDK module registry.

## Why use only one topic is to discover projects?

The `cosmos-sdk` topic is the best way to "discover"  Cosmos Ecosystem projects. The reason at this time to only support this topic is because of the desire that users who will be using this tool are interested in the Cosmos Ecosystem projects and want to find more about activity among the projects (e.g., last updated) or popularity (number of stars) for example.

## Why a rate limit error is returned when running a command?

The tools leverage the unauthenticated calls to the Github API. Github has some limits on how many times calls to its API without authentication (e.g., an API key) can be made. Not enforcing authentication is the best approach for security reasons since it requires less configuration and does not have significant security implications. The same information returned by the Github API is available through the Github website search functionality.