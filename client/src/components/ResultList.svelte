<script lang="ts">
	export let className: string | undefined = undefined;
	export let artists: string[];
	export let bands: string[];
	export let genres: string[];
	export let selectedItem: string | undefined = undefined;
</script>

<aside class={`${className}`}>
	{#if bands.length > 0}
		<section class="rounded-box bg-base-100 mb-4">
			<div class="px-4 py-2">
				<span class="font-semibold">Bands ({bands.length})</span>
			</div>
			<div class="max-h-[25vh] overflow-y-auto">
				<ul class="menu menu-compact">
					{#each bands as band}
						<li><button on:click={() => (selectedItem = band)}>{band}</button></li>
					{/each}
				</ul>
			</div>
		</section>
	{/if}
	{#if artists.length > 0}
		<section class="rounded-box bg-base-100 mb-4">
			<div class="px-4 py-2">
				<span class="font-semibold">Artists ({artists.length})</span>
			</div>
			<div class="max-h-[25vh] overflow-y-auto">
				<ul class="menu menu-compact">
					{#each artists as artist}
						<li><button on:click={() => (selectedItem = artist)}>{artist}</button></li>
					{/each}
				</ul>
			</div>
		</section>
	{/if}
	{#if genres.length > 0}
		<section class="rounded-box bg-base-100 mb-4">
			<div class="px-4 py-2">
				<span class="font-semibold">Genres ({genres.length})</span>
			</div>
			<div class="max-h-[25vh] overflow-y-auto">
				<ul class="menu menu-compact">
					{#each genres as genres}
						<li><button on:click={() => (selectedItem = genres)}>{genres}</button></li>
					{/each}
				</ul>
			</div>
		</section>
	{/if}
</aside>
