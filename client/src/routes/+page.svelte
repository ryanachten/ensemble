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
		selectedItem,
		isLoading,
		hasError
	} from '../stores';
	import LoadingState from '../components/LoadingState.svelte';
	import ErrorState from '../components/ErrorState.svelte';
	import Logo from '../components/Logo.svelte';

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

	$: showResults = !$isLoading && !$hasError;
</script>

<svelte:head>
	<title>ensemble</title>
	<meta name="description" content="Ensemble" />
	<link
		href="https://fonts.googleapis.com/css2?family=Open+Sans:wght@400;600&family=Teko:wght@700&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

{#if $isLoading}
	<div class="absolute top-0 z-5 flex flex-col justify-center items-center h-screen w-screen">
		<LoadingState />
	</div>
{/if}
{#if $hasError}
	<div class="absolute top-0 z-5 flex flex-col justify-center items-center h-screen w-screen">
		<ErrorState />
	</div>
{/if}

<Graph
	className={`h-screen ${!showResults ? 'opacity-0' : ''}`}
	bind:layoutKey={$layoutKey}
	bind:centerGraph
	bind:selectedId={$selectedItem}
/>

<div class="absolute p-4 top-0 z-10 flex justify-between w-screen h-screen pointer-events-none">
	<div class="flex flex-col bg-base-100 h-fit p-4 pointer-events-auto rounded-lg mr-4">
		<Logo />

		<SearchForm onSubmitForm={updateGraph} {onCenterGraph} />
		{#if $nodePath.length > 0}
			<div class="divider" />
			<NodePath />
		{/if}
	</div>

	{#if showResults}
		<ResultLists />
	{/if}
</div>
