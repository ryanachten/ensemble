<script lang="ts">
	import { onMount } from 'svelte';
	import type { ChartParams, CsvRow } from '../../components/charts';
	import ChecksPassedChart from '../../components/charts/ChecksPassedChart.svelte';
	import DurationAvgChart from '../../components/charts/DurationAvgChart.svelte';
	import RequestsFailedChart from '../../components/charts/RequestsFailedChart.svelte';
	import csv from './results.csv';

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

		initCharts({ labels, results });
	}

	onMount(async () => {
		createCharts(csv as CsvRow[]);
	});
</script>

<svelte:head>
	<title>Ensemble - Stats</title>
	<meta name="description" content="Ensemble" />
</svelte:head>

<div class="flex flex-wrap mx-auto max-w-screen-md">
	<ChecksPassedChart className="w-full p-4" bind:init={initChecksPassed} />
	<DurationAvgChart className="w-full p-4" bind:init={initDegreesOfSeparation} />
	<RequestsFailedChart className="w-full p-4" bind:init={initRequestsFailed} />
</div>
