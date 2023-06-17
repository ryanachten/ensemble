import { theme } from '../../theme';

export type TestMode = 'mutex' | 'sequential' | 'syncmap';
export type Endpoint = 'bands' | 'genres';

export interface CsvRow {
	testName: string;
	degreesOfSeparation: string;
	requestsFailed: string;
	checksPassed: string;
	durationAvg: string;
	dateUtc: string;
	mode: TestMode;
	endpoint: Endpoint;
}

export const colors: Record<TestMode, string> = {
	sequential: theme.accent,
	mutex: theme.primary,
	syncmap: theme.secondary
};

export type Dataset = {
	label: string;
	data: { x: string; y: number }[];
	borderColor: string;
	backgroundColor: string;
};

export type ChartParams = {
	labels: string[];
	results: CsvRow[];
};

export const generateDataset = (
	results: CsvRow[],
	labels: string[],
	opts: {
		columnKey: keyof CsvRow;
		formatValue: (value: string) => number;
	}
): Dataset[] => {
	const datasets: Record<string, Dataset> = {};
	results.forEach((result) => {
		const { testName, endpoint, degreesOfSeparation, mode, dateUtc } = result;
		const column = result[opts.columnKey];
		const point = {
			y: opts.formatValue(column),
			x: dateUtc
		};
		if (datasets[testName]) {
			datasets[testName].data.push(point);
		} else {
			datasets[testName] = {
				label: `${mode ? mode : 'syncmap'} | ${endpoint} | ${degreesOfSeparation}°`,
				borderColor: colors[mode] ?? colors.syncmap,
				backgroundColor: colors[mode] ?? colors.syncmap,
				data: [point]
			};
		}
	});

	const response: Dataset[] = [];
	for (const key in datasets) {
		response.push(datasets[key]);
	}
	return response.sort((a, b) => (a.label > b.label ? 1 : -1));
};
