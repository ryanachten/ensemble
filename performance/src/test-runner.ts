import { Endpoint } from "./base-test";

const { exec } = require("child_process");
const converter = require("json-2-csv");
const fs = require("fs");
const path = require("path");

enum GraphMode {
  SYNC = "sync",
  NON_SYNC = "insync",
  MUTEX = "mutex",
}

interface DegreeConfig {
  vus: number;
  iterations: number;
  degreesOfSeparation: number;
  nodeCount: number;
  edgeCount: number;
  endpoint: Endpoint;
}

const modes = [GraphMode.SYNC, GraphMode.MUTEX, GraphMode.NON_SYNC];
const degrees: DegreeConfig[] = [
  {
    endpoint: "bands",
    vus: 2,
    iterations: 5,
    degreesOfSeparation: 1,
    nodeCount: 30,
    edgeCount: 34,
  },
  {
    endpoint: "bands",
    vus: 1,
    iterations: 10,
    degreesOfSeparation: 3,
    nodeCount: 125,
    edgeCount: 167,
  },
  {
    endpoint: "bands",
    vus: 1,
    iterations: 5,
    degreesOfSeparation: 5,
    nodeCount: 641,
    edgeCount: 951,
  },
  {
    endpoint: "genres",
    vus: 1,
    iterations: 1,
    degreesOfSeparation: 1,
    nodeCount: 211,
    edgeCount: 342,
  },
  {
    endpoint: "genres",
    vus: 1,
    iterations: 3,
    degreesOfSeparation: 2,
    nodeCount: 693,
    edgeCount: 2012,
  },
];

interface TestConfig extends DegreeConfig {
  mode: GraphMode;
  dateUtc: string;
}

interface TestResult {
  testName: string;
  dateUtc: string;
  endpoint: Endpoint;
  mode: GraphMode;
  degreesOfSeparation: number;
  vus: number;
  iterations: number;
  checksPassed: string | number | null;
  requestsFailed: string | number | null;
  durationAvg: string | number | null;
}

/**
 * Extracts match from string output or returns null
 */
const getRegexMatch = (regex: RegExp, str: string): string | null => {
  const res = regex.exec(str);
  if (res === null || res.length < 2) return null;
  return res[1];
};

/**
 * Executes K6 test using provided config
 * @param config test execution config
 * @returns promise containing test result
 */
const executeK6Test = ({
  mode,
  dateUtc,
  vus,
  iterations,
  degreesOfSeparation,
  nodeCount,
  edgeCount,
  endpoint,
}: TestConfig): Promise<TestResult> => {
  const testName = `endpoint-${endpoint}-degreesOfSeparation-${degreesOfSeparation}-mode-${
    mode ? mode : "sync"
  }`;
  return new Promise((resolve, reject) => {
    exec(
      `k6 run --vus=${vus} --iterations=${iterations} -e endpoint=${endpoint} -e degreesOfSeparation=${degreesOfSeparation} -e mode=${mode} -e nodeCount=${nodeCount} -e edgeCount=${edgeCount} dist/get-band-test.js`,
      (error: Error, stdout: string, stderr: string) => {
        if (error) {
          console.error(`Error running test ${testName}: ${error.message}`);
          reject(error);
        }
        if (stderr) {
          console.error(`stderr: ${stderr}`);
        }
        const checksPassed = getRegexMatch(
          RegExp(/checks[\.]+\:.([\d,.]*%)/),
          stdout
        );
        const durationAvg = getRegexMatch(
          RegExp(/http_req_duration[\.]+\:.avg=([\d,.]*[\w]+)/),
          stdout
        );
        const requestsFailed = getRegexMatch(
          RegExp(/http_req_failed[\.]+\:.([\d,.]*%)/),
          stdout
        );

        resolve({
          testName,
          dateUtc,
          endpoint,
          mode,
          degreesOfSeparation,
          vus,
          iterations,
          checksPassed,
          durationAvg,
          requestsFailed,
        });
      }
    );
  });
};

/**
 * Runs all the test permutations based on modes and degrees of separation
 */
const runAllTests = async () => {
  const dateUtc = new Date(Date.now()).toUTCString();
  const results: TestResult[] = [];
  const modeExecutions = modes.map(async (mode) => {
    const degreeExecutions = degrees.map(async (config) => {
      const result = await executeK6Test({
        mode,
        dateUtc,
        ...config,
      });
      results.push(result);
    });
    await Promise.all(degreeExecutions);
  });
  await Promise.all(modeExecutions);

  writeToFile(results);
};

/**
 * Writes results to CSV. Appends new results to existing results
 * @param newResults results to append
 */
const writeToFile = async (newResults: TestResult[]) => {
  const filePath = path.join(
    __dirname,
    "../../client/src/routes/stats/results.csv"
  );
  const fileContent = fs.readFileSync(filePath, {
    encoding: "utf8",
    flag: "r",
  });

  const existingResults: TestResult[] = await converter.csv2json(fileContent, {
    trimHeaderFields: true,
    trimFieldValues: true,
  });

  const updatedResults = existingResults.concat(newResults);

  const csv = await converter.json2csv(updatedResults);
  fs.writeFile(filePath, csv, "utf8", (error: Error) => {
    if (error) {
      console.error("Error saving file", error);
    } else {
      console.log("Saved file", filePath);
    }
  });
};

runAllTests();
