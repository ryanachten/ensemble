<script lang="ts">
	import { LayoutKeys } from '../graph/layout';

	export let className: string | undefined = undefined;
	export let layoutKey: LayoutKeys;
	export let bandName: string;
	export let degreesOfSeparation: number;
	export let onSubmitForm: () => void;
</script>

<form class={`${className}`} on:submit={onSubmitForm}>
	<input class="input input-bordered" type="text" bind:value={bandName} />
	<input class="input input-bordered" type="number" bind:value={degreesOfSeparation} />
	<select class="select" name="layout options" id="layoutOptions" bind:value={layoutKey}>
		<option value={LayoutKeys.COSE}>Cose</option>
		<option value={LayoutKeys.BREADTH_FIRST}>Breadth-first</option>
	</select>
	<button class="btn btn-primary" type="submit">Search!</button>
</form>
