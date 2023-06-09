![ensemble splash](./docs/ensemble_splash.png)

Web application visualising connections between bands, artists and genres.

![ensemble screenshot](./docs/ensemble_screenshot_rollingstones.jpg)

## Overview

Ensemble uses Wikipedia as a data source by scraping metadata for a given query and recursively visiting relevant links. These relationships are dynamically built up as an adjacency list before being formatted for client consumption once a given degree-of-separation is met.

![ensemble sequence](./docs/ensemble_sequence.png)

### Stack

Ensemble is built using the following technologies:

**Client**

- [Svelte + Sveltekit](https://svelte.dev/) - Client framework
- [Tailwind](https://tailwindcss.com/) + [DaisyUI](https://daisyui.com/) - Styling
- [Cytoscape](https://js.cytoscape.org/) - Graph rendering and searching

**API**

- [Go](https://go.dev/) + [Gin](https://gin-gonic.com/) - API framework
- [Colly](https://go-colly.org/) - Web scraping

**Performance testing**

- [Node](https://nodejs.org/en) - tooling
- [K6](https://k6.io/) - performance testing

![ensemble screenshot](./docs/ensemble_screenshot_blackflag.jpg)
![ensemble screenshot](./docs/ensemble_screenshot_royorbison.jpg)

## Graph relationships

In the band or artist graph, the following relationships are supported in addition to their corresponding data source in the Wikipedia website. The genre graph is a little different, where only genre nodes are currently included.

![ensemble band nodes](./docs/ensemble_band_nodes.png)

## Degrees of Separation

To prevent infinite recursion, we need a hard limit how how many times to recursion cycles we'll complete when building the graph. In the Ensemble, this is limit is referred to as _degrees of separation_ from the original target.

![ensemble degrees of separation](./docs/ensemble_dos.png)

## Pathfinding

Currently, Ensemble uses Cytoscape's implementation of the [Floyd-Warshall](https://js.cytoscape.org/#eles.floydWarshall) search algorithm for finding the shortest paths between two nodes. Genre nodes are weighted slightly lower than band or artist nodes to discourage routing via genres for more interesting paths.

![ensemble pathfinding](./docs/ensemble_screenshot_wutang.jpg)
![ensemble pathfinding](./docs/ensemble_screenshot_sleep.jpg)

<!-- TODO: add performance test docs -->

## Further information

- [Running ensemble client](./client/README.md)
- [Running ensemble API](./api/README.md)
- [Running ensemble performance tests](./performance/README.md)
