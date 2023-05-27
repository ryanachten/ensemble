<script lang="ts">
	import { nodePath, confirmedNodePath } from '../stores';
	import { nodeColourMap } from '../theme';

	const confirmNodePath = () => confirmedNodePath.set($nodePath);
	const resetNodePath = () => {
		nodePath.set([]);
		confirmedNodePath.set([]);
	};
</script>

<form>
	<ul class="steps mb-4 w-full">
		{#each $nodePath as node}
			<li class={`step step-${nodeColourMap[node.type]}`}>{node.label}</li>
		{/each}
	</ul>
	<div class="grid grid-cols-2 gap-4">
		{#if $nodePath.length > 1}
			<button class="btn btn-primary btn-sm" type="submit" on:click={confirmNodePath}
				>Find path</button
			>
		{/if}
		<button class="btn btn-sm" type="reset" on:click={resetNodePath}>Reset</button>
	</div>
</form>
