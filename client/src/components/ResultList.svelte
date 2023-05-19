<script lang="ts">
	export let items: string[];
	export let title: string;
	export let selectedItem: string | undefined;
</script>

<section class="rounded-box bg-base-100 mb-4">
	<div class="px-4 py-2">
		<span class="font-semibold">{title} ({items.length})</span>
	</div>
	<div class="max-h-[25vh] overflow-y-auto">
		<ul class="menu menu-compact">
			{#each items as item}
				<li>
					<button class:active={selectedItem === item} on:click={() => (selectedItem = item)}
						>{item}</button
					>
				</li>
			{/each}
		</ul>
	</div>
</section>
