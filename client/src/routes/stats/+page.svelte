<script lang="ts">
	import { onMount } from 'svelte';
	import { csv } from 'd3';
	import type { ChartParams, CsvRow } from '../../components/charts';
	import ChecksPassedChart from '../../components/charts/ChecksPassedChart.svelte';
	import DurationAvgChart from '../../components/charts/DurationAvgChart.svelte';
	import RequestsFailedChart from '../../components/charts/RequestsFailedChart.svelte';

	let initChecksPassed: (params: ChartParams) => void;
	let initDegreesOfSeparation: (params: ChartParams) => void;
	let initRequestsFailed: (params: ChartParams) => void;
	let initCharts = (params: ChartParams) => {
		initChecksPassed(params);
		initDegreesOfSeparation(params);
		initRequestsFailed(params);
	};

	function createCharts(results: CsvRow[]) {
		const dates = new Set<string>();
		results.forEach((d) => dates.add(d.dateUtc));
		const labels = Array.from(dates);
		const formattedLabels = labels.map((d) => new Date(d).toDateString());

		initCharts({ labels, formattedLabels, results });
	}

	onMount(async () => {
		const data = await csv<keyof CsvRow>(
			'https://raw.githubusercontent.com/ryanachten/ensemble/main/performance/output/results.csv'
		);

		createCharts(data as CsvRow[]);
	});
</script>

<svelte:head>
	<title>Ensemble - Stats</title>
	<meta name="description" content="Ensemble" />
</svelte:head>

<div class="flex flex-wrap mx-auto max-w-screen-sm">
	<ChecksPassedChart className="w-full p-4" bind:init={initChecksPassed} />
	<DurationAvgChart className="w-full p-4" bind:init={initDegreesOfSeparation} />
	<RequestsFailedChart className="w-full p-4" bind:init={initRequestsFailed} />
</div>
