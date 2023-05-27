<script lang="ts">
	import { onMount } from 'svelte';
	import SearchForm from '../components/SearchForm.svelte';
	import ResultLists from '../components/ResultLists.svelte';
	import Graph from '../components/Graph.svelte';
	import { requestGraph } from '../api';
	import NodePath from '../components/NodePath.svelte';
	import {
		degreesOfSeparation,
		nodePath,
		resource,
		searchTerm,
		layoutKey,
		selectedItem
	} from '../stores';

	const updateGraph = () =>
		requestGraph({
			resource: $resource,
			name: $searchTerm,
			degreesOfSeparation: $degreesOfSeparation
		});

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
	bind:layoutKey={$layoutKey}
	bind:centerGraph
	bind:selectedId={$selectedItem}
/>

<div class="absolute p-4 top-0 z-10 flex justify-between w-screen h-screen pointer-events-none">
	<div class="flex flex-col bg-base-100 h-fit p-4 pointer-events-auto rounded-lg mr-4">
		<SearchForm onSubmitForm={updateGraph} {onCenterGraph} />
		{#if $nodePath.length > 0}
			<div class="divider" />
			<NodePath />
		{/if}
	</div>

	<ResultLists className="pointer-events-auto" />
</div>
