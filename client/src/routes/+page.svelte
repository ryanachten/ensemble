<script lang="ts">
	import { onMount } from 'svelte';
	import graph from '../graph';
	import type { ElementsDefinition } from 'cytoscape';

	let container: HTMLDivElement;

	onMount(async () => {
		const res = await fetch('http://localhost:8080/bands?name=Black Flag');
		const data = (await res.json()) as ElementsDefinition;
		graph(data, container);
	});
</script>

<svelte:head>
	<title>Ensemble</title>
	<meta name="description" content="Ensemble" />
</svelte:head>

<div bind:this={container} class="graph-container" />

<style>
	.graph-container {
		min-height: 100vh;
	}
</style>
