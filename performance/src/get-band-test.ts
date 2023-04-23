import { BandTest, GraphMode } from "./base-test";

export default () => {
  const degreesOfSeparation = parseInt(__ENV["degreesOfSeparation"]);
  const mode = __ENV["mode"];
  const nodeCount = parseInt(__ENV["nodeCount"]);
  const edgeCount = parseInt(__ENV["edgeCount"]);

  if (Boolean(mode) && mode != GraphMode.MUTEX && mode != GraphMode.NON_SYNC)
    return console.error("mode invalid", mode);
  if (!degreesOfSeparation) return console.error("degreesOfSeparation not set");
  if (!nodeCount) return console.error("nodeCount not set");
  if (!edgeCount) return console.error("edgeCount not set");

  BandTest({
    mode: mode as GraphMode,
    degreesOfSeparation,
    expectedResponse: {
      nodeCount,
      edgeCount,
    },
  });
};
