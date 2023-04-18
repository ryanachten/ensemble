<script lang="ts">
	import { onMount } from 'svelte';
	import type { Graph } from '../models/api';
	import graph from '../graph';

	let container: HTMLDivElement;

	onMount(async () => {
		const res = await fetch('http://localhost:8080/bands?name=Black Flag');
		const data = (await res.json()) as Graph;
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
