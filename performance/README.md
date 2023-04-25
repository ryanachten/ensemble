# Performance Testing

Testing graph implementation performance and reliability using K6. Based on the [K6 TypeScript template](https://github.com/grafana/k6-template-typescript)

## Prerequisites

- Node v18
- Yarn
- [K6 CLI](https://k6.io/docs/get-started/installation/)

## Usage

- Install dependencies `yarn install`
- Build test scripts `yarn build` (append `--watch` to watch for changes)
- Run test manually `k6 run --vus=1 --iterations=10 -e degreesOfSeparation=3 -e mode=mutex -e nodeCount=101 -e edgeCount=132 dist/get-band-test.js`
- Or use custom test runner for running all tests and saving results to CSV `yarn test`
