# Frequently Asked Questions

- [Why my Cosmos project is not showing on the results?](#why-my-cosmos-project-is-not-showing-on-the-results)
- [Are these projects displayed official Cosmos SDK projects?](#are-these-projects-displayed-official-cosmos-sdk-projects)
- [Why only one topic is used to discover projects?](#why-only-one-topic-is-used-to-discover-projects)
- [Why I am getting a rate limit error when running a command?](#why-i-am-getting-a-rate-limit-error-when-running-a-command)

## Why my Cosmos project is not showing on the results?

Currently the `cosmocope` tool only retrieves Github repositories that have a topic named `cosmos-sdk` associated with it. Please add the topic and run the tool again. It should show up on the next time you run `cosmocope discover projects`. 

## Are these projects displayed official Cosmos SDK projects?

The projects returned are not part of a curated list of Cosmos projects. If someone creates a project that is not based in the `cosmos-sdk` and decides to tag them with a `cosmos-sdk` project there is nothing to prevent that since this is a Github feature available to anyone who has access to Github. 

Please note that the repositories returned might be something other a Cosmos SDK chain. A lot of developers use the `cosmic-sdk` topic to tag their projects as part of the Cosmos ecosystem, so the list returned might include projects that are for wallets or client libraries for example and not only Cosmos chains.

In the future, there might be better ways to list the real Cosmos SDK repositories maybe based on a vetted or curated list of repositories. 

The intention of this tool is only to allow discoverability of new projects that might become part of the Cosmos Ecosystem. For example, if you manage or host a list of Cosmos projects, you can use this tools to find about new projects and then after some due diligence and assessment you could add a new project to your list.

For an official list of Cosmos SDK projects please check the official [Cosmos Network Ecosystem](https://cosmos.network/ecosystem) website.

For an official list of modules, please check [Atlas](https://atlas.cosmos.network) which is the official Cosmos SDK module registry.

## Why only one topic is used to discover projects?

The `cosmos-sdk` topic is the best way to "discover"  Cosmos Ecosystem projects. The reason at this time to only support this topic is because of the desire that users who will be using this tool are interested in the Cosmos Ecosystem projects and want to find more about activity among the projects (e.g. last updated) or popularity (number of stars) for example.

## Why I am getting a rate limit error when running a command?

The tools leverage the unauthenticated calls to the Github API. Github has some limits on how many times their API can be called with authentication (e.g. an API key). For security reasons, this might be the best approach this time since it requires less configuration and doesn't have major security implications. The information returned by the API could be also discovered through the Github website search functionality.

