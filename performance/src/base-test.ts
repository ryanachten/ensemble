import { group, check, sleep } from "k6";
import http, { ResponseBody } from "k6/http";
import { Options } from "k6/options";

// Only define properties relevant to the test
type GraphMetadata = {
  nodeCount: number;
  edgeCount: number;
};

type GraphResponse = ResponseBody & GraphMetadata;

export type Endpoint = "bands" | "genres";

interface TestConfig {
  mode: GraphMode;
  degreesOfSeparation: number;
  expectedResponse: GraphMetadata;
  endpoint: Endpoint;
}

export enum GraphMode {
  SYNC = "sync",
  NON_SYNC = "insync",
  MUTEX = "mutex",
}

export let baseOptions: Options = {
  iterations: 10,
};

const BASE_URL = "http://localhost:8080";

export const BandTest = ({
  mode,
  degreesOfSeparation,
  expectedResponse,
  endpoint,
}: TestConfig) =>
  group(`Request using ${mode}`, () => {
    let name = endpoint === "bands" ? "Melvins" : "Hardcore%20punk";
    const res = http.get(
      `${BASE_URL}/${endpoint}?name=${name}&degreesOfSeparation=${degreesOfSeparation}&mode=${mode}`
    );
    check(res, {
      "status is 200": () => res.status === 200,
    });
    check(res, {
      "has correct node count": () => {
        const graph = res.json() as GraphResponse;
        const hasCorrectCount = graph.nodeCount === expectedResponse.nodeCount;
        if (!hasCorrectCount) {
          console.log(
            `Expected ${expectedResponse.nodeCount} nodes, but received`,
            graph.nodeCount
          );
        }
        return hasCorrectCount;
      },
    });
    check(res, {
      "has correct edge count": () => {
        const graph = res.json() as GraphResponse;
        const hasCorrectCount = graph.edgeCount === expectedResponse.edgeCount;
        if (!hasCorrectCount) {
          console.log(
            `Expected ${expectedResponse.edgeCount} edges, but received`,
            graph.edgeCount
          );
        }
        return hasCorrectCount;
      },
    });
    sleep(2);
  });
