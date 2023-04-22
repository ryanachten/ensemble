<script lang="ts">
	import { onMount } from 'svelte';
	import type { ElementsDefinition } from 'cytoscape';
	import { LayoutKeys } from '../graph/layout';
	import SearchForm from '../components/SearchForm.svelte';
	import ResultList from '../components/ResultList.svelte';
	import Graph from '../components/Graph.svelte';
	import { requestBandGraph } from '../api/bands';

	let bandName = 'Black Flag';
	let degreesOfSeparation = 3;
	let layoutKey = LayoutKeys.COSE;
	let artists: string[] = [];
	let bands: string[] = [];
	let data: ElementsDefinition | null = null;
	let selectedItem: string | undefined;

	const updateGraph = async () => {
		const {
			data: updatedData,
			bands: updatedBands,
			artists: updatedArtists
		} = await requestBandGraph(bandName, degreesOfSeparation);
		data = updatedData;
		bands = updatedBands;
		artists = updatedArtists;
	};

	onMount(async () => {
		updateGraph();
	});

	let centerGraph: () => void;
	const onCenterGraph = () => centerGraph();
</script>

<svelte:head>
	<title>Ensemble</title>
	<meta name="description" content="Ensemble" />
</svelte:head>

<Graph
	className="h-screen"
	bind:layoutKey
	bind:data
	bind:centerGraph
	bind:selectedId={selectedItem}
/>

<SearchForm
	className="absolute p-4 top-0 z-10"
	bind:layoutKey
	bind:bandName
	bind:degreesOfSeparation
	onSubmitForm={updateGraph}
	{onCenterGraph}
/>

<ResultList className="absolute p-4 top-0 right-0 z-10" bind:bands bind:artists bind:selectedItem />
