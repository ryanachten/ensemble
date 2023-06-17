import { Endpoint } from "./base-test";
import config, { DegreeConfig } from "./config";

const { exec } = require("child_process");
const converter = require("json-2-csv");
const fs = require("fs");
const path = require("path");
const { program } = require("commander");

interface TestConfig extends DegreeConfig {
  mode: GraphMode;
  dateUtc: string;
  degreesOfSeparation: number;
  endpoint: Endpoint;
}

export enum GraphMode {
  SYNC = "sync",
  SEQUENTIAL = "sequential",
  MUTEX = "mutex",
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

program
  .option(
    "-dos, --degreesOfSeparation <items>",
    "degrees of separation",
    (value: string) => value.split(" ").map((x) => parseInt(x))
  )
  .option("-e, --endpoints <items>", "endpoints", (value: string) =>
    value.split(" ")
  )
  .option("-m, --modes <items>", "modes", (value: string) => value.split(" "));

program.parse();

const {
  degreesOfSeparation = [1, 2, 3, 5],
  endpoints = ["bands", "genres"],
  modes = [GraphMode.SYNC, GraphMode.MUTEX, GraphMode.SEQUENTIAL],
}: {
  degreesOfSeparation: number[];
  endpoints: Endpoint[];
  modes: GraphMode[];
} = program.opts();

console.log(
  "Running tests with the following settings:",
  "\ndegreesOfSeparation:",
  degreesOfSeparation,
  "\nendpoints:",
  endpoints,
  "\nmodes:",
  modes
);

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
 * Runs all the test permutations based on endpoints, modes and degrees of separation
 */
const runAllTests = async () => {
  const dateUtc = new Date(Date.now()).toUTCString();
  const tests: TestConfig[] = [];

  endpoints.forEach((endpoint) =>
    degreesOfSeparation.forEach((dos) =>
      modes.forEach((mode) => {
        const degreeConfig = config[endpoint][dos];
        if (degreeConfig) {
          tests.push({
            ...degreeConfig,
            mode,
            dateUtc,
            endpoint,
            degreesOfSeparation: dos,
          });
        }
      })
    )
  );

  const executions = tests.map(async (test) => await executeK6Test(test));
  const results = await Promise.all(executions);

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
