<script lang="ts">
	import ResultList from './ResultList.svelte';

	export let className: string | undefined = undefined;
	export let artists: string[];
	export let bands: string[];
	export let genres: string[];
	export let selectedItem: string | undefined = undefined;
</script>

<aside class={`${className}`}>
	{#if bands.length > 0}
		<ResultList title="Bands" items={bands} bind:selectedItem />
	{/if}
	{#if artists.length > 0}
		<ResultList title="Artists" items={artists} bind:selectedItem />
	{/if}
	{#if genres.length > 0}
		<ResultList title="Genres" items={genres} bind:selectedItem />
	{/if}
</aside>
