<script lang="ts">
	import { requestGraph, type Resource } from '../api';
	import type { NodeMetadata, NodeType } from '../interfaces';
	import { degreesOfSeparation, nodePath, resource, searchTerm } from '../stores';
	import { nodeColourMap } from '../theme';
	import Icon from './Icon.svelte';

	export let data: NodeMetadata;

	const addItem = () => nodePath.update((p) => [...p, data]);
	const removeItem = () => nodePath.update((p) => p.filter((x) => x.id !== data.id));
	const nodeTypeToResource = (nodeType: NodeType): Resource => {
		const resource = `${nodeType}s`;
		return resource as Resource;
	};
	const searchItem = () => {
		const dataResource = nodeTypeToResource(data.type);
		searchTerm.set(data.label);
		resource.set(dataResource);
		requestGraph({
			name: data.label,
			resource: dataResource,
			degreesOfSeparation: $degreesOfSeparation
		});
	};

	const wikipediaUrl = `https://en.wikipedia.org/w/index.php?title=${encodeURIComponent(
		data.label
	)}`;
</script>

<div class="card bg-base-100 shadow-xl z-20">
	<div class="card-body">
		<h2 class="card-title">{data.label}</h2>
		<span class={`badge badge-${nodeColourMap[data.type]}`}>{data.type}</span>
		<div>
			<span class="mr-2">Search graph</span>
			<button class="btn btn-circle btn-sm btn-primary" on:click={searchItem}
				><Icon name="search" /></button
			>
		</div>
		<div>
			<a class="flex items-center mr-2" href={wikipediaUrl} target="_blank"
				>View on Wikipedia <Icon name="external-link" className="ml-2 inline" /></a
			>
		</div>
		<div class="divider" />
		<div>
			{#if $nodePath.find((x) => x.id === data.id)}
				<span class="mr-2">Remove from path</span>
				<button class="btn btn-circle btn-sm btn-warning" on:click={removeItem}
					><Icon name="minus" />
				</button>
			{:else}
				<span class="mr-2">Add to path</span>
				<button class="btn btn-circle btn-sm btn-warning" on:click={addItem}
					><Icon name="plus" /></button
				>
			{/if}
		</div>
	</div>
</div>
