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
	<div class="flex">
		<ul class="steps flex-grow">
			{#each $nodePath as node}
				<li class={`step step-${nodeColourMap[node.type]}`}>{node.label}</li>
			{/each}
		</ul>
		<button class="btn mr-2" type="reset" on:click={resetNodePath}>Reset</button>
		{#if $nodePath.length > 1}
			<button class="btn" type="submit" on:click={confirmNodePath}>Find path</button>
		{/if}
	</div>
</form>
