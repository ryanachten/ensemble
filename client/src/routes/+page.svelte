<script lang="ts">
	import { onMount } from 'svelte';
	import type { ElementsDefinition } from 'cytoscape';
	import { LayoutKeys } from '../graph/layout';
	import SearchForm from '../components/SearchForm.svelte';
	import ResultLists from '../components/ResultLists.svelte';
	import Graph from '../components/Graph.svelte';
	import { requestGraph, type Resource } from '../api';

	let name = 'Black Flag';
	let resource: Resource = 'bands';
	let degreesOfSeparation = 3;
	let layoutKey = LayoutKeys.COSE;
	let artists: string[] = [];
	let bands: string[] = [];
	let genres: string[] = [];
	let data: ElementsDefinition | null = null;
	let selectedItem: string | undefined;

	const updateGraph = async () => {
		const {
			data: updatedData,
			bands: updatedBands,
			artists: updatedArtists,
			genres: updatedGenres
		} = await requestGraph({ resource, name, degreesOfSeparation });
		data = updatedData;
		bands = updatedBands;
		artists = updatedArtists;
		genres = updatedGenres;
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
	bind:name
	bind:resource
	bind:degreesOfSeparation
	onSubmitForm={updateGraph}
	{onCenterGraph}
/>

<ResultLists
	className="absolute p-4 top-0 right-0 z-10"
	bind:bands
	bind:artists
	bind:genres
	bind:selectedItem
/>
