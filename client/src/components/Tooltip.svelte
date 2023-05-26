<script lang="ts">
	import type { NodeMetadata } from '../interfaces';
	import { nodePath } from '../stores';
	import { nodeColourMap } from '../theme';

	export let data: NodeMetadata;
	const addItem = () => nodePath.update((p) => [...p, data]);
	const removeItem = () => nodePath.update((p) => p.filter((x) => x.id !== data.id));
</script>

<div class="card bg-base-100 shadow-xl z-20">
	<div class="card-body">
		<h2 class="card-title">{data.label}</h2>
		<span class={`badge badge-${nodeColourMap[data.type]}`}>{data.type}</span>
		{#if $nodePath.find((x) => x.id === data.id)}
			<button class="btn btn-circle btn-error" on:click={removeItem}>-</button>
		{:else}
			<button class="btn btn-circle btn-success" on:click={addItem}>+</button>
		{/if}
	</div>
</div>
