<script lang="ts">
	import { onMount } from 'svelte';
	import graph from '../graph';
	import type { Core, ElementsDefinition } from 'cytoscape';
	import { LayoutKeys, layouts } from '../graph/layout';
	import { NodeType, type NodeData } from '../models/api';

	let container: HTMLDivElement;
	let bandName = 'Black Flag';
	let degreesOfSeparation = 3;
	let cytoscape: Core | null = null;
	let layoutKey = LayoutKeys.COSE;
	let artists: string[] = [];
	let bands: string[] = [];

	const requestGraph = async () => {
		cytoscape?.destroy();
		const res = await fetch(
			`http://localhost:8080/bands?name=${bandName}&degreesOfSeparation=${degreesOfSeparation}`
		);
		const data = (await res.json()) as ElementsDefinition;
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

		cytoscape = graph({ data, container, layout: layouts[layoutKey] });
	};

	const updateLayout = () => {
		const newLayout = layouts[layoutKey];
		const updatedLayout = cytoscape?.layout(newLayout);
		updatedLayout?.run();
	};

	onMount(async () => {
		requestGraph();
	});

	const onSubmitForm = async () => {
		requestGraph();
	};
</script>

<svelte:head>
	<title>Ensemble</title>
	<meta name="description" content="Ensemble" />
</svelte:head>

<form class="form-container" on:submit={onSubmitForm}>
	<input type="text" bind:value={bandName} />
	<input type="number" bind:value={degreesOfSeparation} />
	<select name="layout options" id="layoutOptions" bind:value={layoutKey} on:change={updateLayout}>
		<option value={LayoutKeys.COSE}>Cose</option>
		<option value={LayoutKeys.BREADTH_FIRST}>Breadth-first</option>
	</select>
	<button type="submit">Search!</button>
</form>
<aside class="list-container">
	<strong>Bands</strong>
	<ul class="list">
		{#each bands as band}
			<li>{band}</li>
		{/each}
	</ul>
	<strong>Artists</strong>
	<ul class="list">
		{#each artists as artist}
			<li>{artist}</li>
		{/each}
	</ul>
</aside>
<div bind:this={container} class="graph-container" />

<style>
	.form-container {
		position: absolute;
		z-index: 100;
	}
	.graph-container {
		height: 100vh;
	}
	.list-container {
		max-height: 100vh;
		overflow-y: auto;
		position: absolute;
		right: 0;
		text-align: right;
		z-index: 100;
	}
	.list {
		list-style: none;
	}
</style>
