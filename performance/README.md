# Performance Testing

Testing graph implementation performance and reliability using K6. Based on the [K6 TypeScript template](https://github.com/grafana/k6-template-typescript)

## Prerequisites

- Node v18
- Yarn
- [K6 CLI](https://k6.io/docs/get-started/installation/)

## Usage

- Install dependencies `yarn install`
- Build test scripts `yarn build` (append `--watch` to watch for changes)
- Run test manually `k6 run --vus=1 --iterations=10 -e endpoint=bands -e degreesOfSeparation=3 -e mode=mutex -e nodeCount=125 -e edgeCount=167 dist/get-band-test.js`
- Or use custom test runner for running all tests and saving results to CSV `yarn test`

## Memory profiling

To profile memory in Go, we can use the `pprof` tool to investigate bottlenecks

- Produce flow chart of memory allocation displayed as interactive SVG `go tool pprof -web http://localhost:8080/debug/pprof/heap`
- Find top memory allocations
  - Enter interactive mode `go tool pprof http://localhost:8080/debug/pprof/heap`
  - Execute `top` command
