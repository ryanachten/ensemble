import { Endpoint } from "./base-test";

export interface DegreeConfig {
  vus: number;
  iterations: number;
  nodeCount: number;
  edgeCount: number;
}

const config: Record<Endpoint, Record<number, DegreeConfig>> = {
  bands: {
    1: {
      vus: 2,
      iterations: 5,
      nodeCount: 30,
      edgeCount: 34,
    },
    3: {
      vus: 1,
      iterations: 10,
      nodeCount: 125,
      edgeCount: 167,
    },
    5: {
      vus: 1,
      iterations: 5,
      nodeCount: 641,
      edgeCount: 951,
    },
  },
  genres: {
    1: {
      vus: 1,
      iterations: 1,
      nodeCount: 211,
      edgeCount: 342,
    },
    2: {
      vus: 1,
      iterations: 3,
      nodeCount: 693,
      edgeCount: 2012,
    },
  },
};

export default config;
