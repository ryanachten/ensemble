<script lang="ts">
	import { onMount } from 'svelte';
	import type { ElementsDefinition } from 'cytoscape';
	import { LayoutKeys } from '../graph/layout';
	import { NodeType, type NodeData } from '../models/api';
	import SearchForm from '../components/SearchForm.svelte';
	import ResultList from '../components/ResultList.svelte';
	import Graph from '../components/Graph.svelte';

	let bandName = 'Black Flag';
	let degreesOfSeparation = 3;
	let layoutKey = LayoutKeys.COSE;
	let artists: string[] = [];
	let bands: string[] = [];
	let data: ElementsDefinition | null = null;
	let selectedItem: string | undefined;

	const requestGraph = async () => {
		const res = await fetch(
			`http://localhost:8080/bands?name=${bandName}&degreesOfSeparation=${degreesOfSeparation}`
		);
		data = (await res.json()) as ElementsDefinition;
		const updatedBands: string[] = [];
		const updatedArtists: string[] = [];
		data.nodes.forEach((node) => {
			const data = node.data as NodeData;
			if (data.type === NodeType.ARTIST && data.id) {
				updatedArtists.push(data.id);
			}
			if (data.type === NodeType.BAND && data.id) {
				updatedBands.push(data.id);
			}
		});
		bands = updatedBands.sort();
		artists = updatedArtists.sort();
	};

	onMount(async () => {
		requestGraph();
	});
</script>

<svelte:head>
	<title>Ensemble</title>
	<meta name="description" content="Ensemble" />
</svelte:head>

<Graph className="h-screen" bind:layoutKey bind:data />

<SearchForm
	className="absolute p-4 top-0 z-10"
	bind:layoutKey
	bind:bandName
	bind:degreesOfSeparation
	onSubmitForm={() => requestGraph()}
/>

<ResultList
	className="absolute max-h-screen overflow-y-auto p-4 top-0 right-0 z-10"
	bind:bands
	bind:artists
	bind:selectedItem
/>
