const { exec } = require("child_process");
const converter = require("json-2-csv");
const fs = require("fs");
const path = require("path");

enum GraphMode {
  SYNC = "",
  NON_SYNC = "nonsync",
  MUTEX = "mutex",
}

const modes = [GraphMode.SYNC, GraphMode.MUTEX]; // GraphMode.NON_SYNC no longer being tested
const degrees = [
  {
    vus: 2,
    iterations: 5,
    degreesOfSeparation: 1,
    nodeCount: 25,
    edgeCount: 29,
  },
  {
    vus: 1,
    iterations: 10,
    degreesOfSeparation: 3,
    nodeCount: 101,
    edgeCount: 132,
  },
  {
    vus: 1,
    iterations: 5,
    degreesOfSeparation: 5,
    nodeCount: 557,
    edgeCount: 785,
  },
];

interface TestResult {
  testName: string;
  dateUtc: string;
  mode: GraphMode;
  degreesOfSeparation: number;
  vus: number;
  iterations: number;
  checksPassed: string | number | null;
  requestsFailed: string | number | null;
  durationAvg: string | number | null;
}

interface TestConfig {
  mode: GraphMode;
  dateUtc: string;
  vus: number;
  iterations: number;
  degreesOfSeparation: number;
  nodeCount: number;
  edgeCount: number;
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
}: TestConfig): Promise<TestResult> => {
  const testName = `degreesOfSeparation-${degreesOfSeparation}-mode-${
    mode ? mode : "sync"
  }`;
  return new Promise((resolve, reject) => {
    exec(
      `k6 run --vus=${vus} --iterations=${iterations} -e degreesOfSeparation=${degreesOfSeparation} -e mode=${mode} -e nodeCount=${nodeCount} -e edgeCount=${edgeCount} dist/get-band-test.js`,
      (error: Error, stdout: string, stderr: string) => {
        if (error) {
          console.error(`Error running test ${testName}: ${error.message}`);
          reject(error);
        }
        if (stderr) {
          console.error(`stderr: ${stderr}`);
        }
        // console.log(`stdout: ${stdout}`);
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
  const testRuns: Promise<TestResult>[] = [];
  modes.forEach((mode) => {
    degrees.forEach((config) => {
      const promise = executeK6Test({
        mode,
        dateUtc,
        ...config,
      });
      testRuns.push(promise);
    });
  });

  const results = await Promise.all(testRuns);
  writeToFile(results);
};

/**
 * Writes results to CSV. Appends new results to existing results
 * @param newResults results to append
 */
const writeToFile = async (newResults: TestResult[]) => {
  const filePath = path.join(__dirname, "../output/results.csv");
  const fileContent = fs.readFileSync(filePath, {
    encoding: "utf8",
    flag: "r",
  });

  const existingResults: TestResult[] = await converter.csv2json(fileContent);
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
