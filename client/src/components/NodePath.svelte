<script lang="ts">
	import { nodePath, confirmedNodePath } from '../stores';
	import Icon from './Icon.svelte';

	const confirmNodePath = () => confirmedNodePath.set($nodePath);
	const resetNodePath = () => {
		nodePath.set([]);
		confirmedNodePath.set([]);
	};
</script>

<form>
	<ul class="steps mb-4 w-full">
		{#each $nodePath as node}
			<li class="step step-warning">{node.label}</li>
		{/each}
	</ul>
	<div class="grid grid-cols-2 gap-4">
		{#if $nodePath.length > 1}
			<button class="btn btn-warning btn-sm" type="submit" on:click={confirmNodePath}
				><Icon className="mr-2" name="compass" />Find path</button
			>
		{/if}
		<button class="btn btn-sm" type="reset" on:click={resetNodePath}
			><Icon className="mr-2" name="rotate-ccw" />Reset</button
		>
	</div>
</form>
