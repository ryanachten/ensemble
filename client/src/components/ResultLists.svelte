<script lang="ts">
	import { artists, bands, genres } from '../stores';
	import ResultList from './ResultList.svelte';

	export let className: string | undefined = undefined;
</script>

<aside class={`${className}`}>
	{#if $bands.length > 0}
		<ResultList title="Bands" items={$bands} />
	{/if}
	{#if $artists.length > 0}
		<ResultList title="Artists" items={$artists} />
	{/if}
	{#if $genres.length > 0}
		<ResultList title="Genres" items={$genres} />
	{/if}
</aside>
