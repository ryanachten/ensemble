<script lang="ts">
	import type { NodeData } from '../interfaces';
	import { selectedItem } from '../stores';

	export let items: NodeData[];
	export let title: string;
	export let titleClassName: string;
</script>

<section class="rounded-box bg-base-100 mb-4 pointer-events-auto shadow-md">
	<div class="px-4 pt-4 pb-2">
		<span class={`font-semibold text-sm ${titleClassName} px-2 py-1 rounded-md`}
			>{title} ({items.length})</span
		>
	</div>
	<div class="max-h-[20vh] overflow-y-auto">
		<ul class="menu menu-compact">
			{#each items as item}
				<li>
					<button
						class:active={$selectedItem === item.id}
						on:click={() => selectedItem.set(item.id ?? '')}>{item.label}</button
					>
				</li>
			{/each}
		</ul>
	</div>
</section>
