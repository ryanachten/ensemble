import { BandTest, Endpoint, GraphMode } from "./base-test";

export default () => {
  const degreesOfSeparation = parseInt(__ENV["degreesOfSeparation"]);
  const mode = __ENV["mode"] as GraphMode;
  const endpoint = __ENV["endpoint"] as Endpoint;
  const nodeCount = parseInt(__ENV["nodeCount"]);
  const edgeCount = parseInt(__ENV["edgeCount"]);

  const validModes = [GraphMode.SYNC, GraphMode.MUTEX, GraphMode.SEQUENTIAL];
  if (Boolean(mode) && !validModes.includes(mode as GraphMode))
    return console.error("mode invalid", mode);

  const validEndpoints: Endpoint[] = ["genres", "bands"];
  if (!validEndpoints.includes(endpoint)) {
    return console.error("endpoint invalid", endpoint);
  }

  if (!degreesOfSeparation) return console.error("degreesOfSeparation not set");
  if (!nodeCount) return console.error("nodeCount not set");
  if (!edgeCount) return console.error("edgeCount not set");

  BandTest({
    mode,
    degreesOfSeparation,
    expectedResponse: {
      nodeCount,
      edgeCount,
    },
    endpoint,
  });
};
