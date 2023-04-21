<script lang="ts">
	export let className: string | undefined = undefined;
	export let bands: string[];
	export let artists: string[];
	export let selectedItem: string | undefined = undefined;
</script>

<aside class={`${className} text-right`}>
	{#if bands.length > 0}
		<strong class="block mb-2">Bands</strong>
		<ul class="mb-4">
			{#each bands as band}
				<li><button on:click={() => (selectedItem = band)}>{band}</button></li>
			{/each}
		</ul>
	{/if}
	{#if artists.length > 0}
		<strong class="block mb-2">Artists</strong>
		<ul>
			{#each artists as artist}
				<li><button on:click={() => (selectedItem = artist)}>{artist}</button></li>
			{/each}
		</ul>
	{/if}
</aside>
