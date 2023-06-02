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

	$: isInPath = $nodePath.find((x) => x.id === data.id);
</script>

<div class="card bg-base-100 shadow-xl z-20">
	<div class="card-body">
		<h2 class="card-title">{data.label}</h2>
		<span class={`badge badge-${nodeColourMap[data.type]}`}>{data.type}</span>
		<div class="grid grid-cols-[max-content] gap-4 items-center">
			<span>Search graph</span>
			<button class="btn btn-circle btn-sm btn-primary" on:click={searchItem}
				><Icon name="search" /></button
			>
			<span>View on Wikipedia</span>
			<a class="btn btn-circle btn-sm" href={wikipediaUrl} target="_blank">
				<Icon name="external-link" className="inline" /></a
			>
			<div class="divider m-0 col-span-2" />
			<span>{isInPath ? 'Remove from path' : 'Add to path'}</span>
			<button
				class="swap btn btn-circle btn-sm btn-warning"
				class:swap-active={isInPath}
				on:click={isInPath ? removeItem : addItem}
			>
				<Icon className="swap-on" name="minus" />
				<Icon className="swap-off" name="plus" />
			</button>
		</div>
	</div>
</div>
