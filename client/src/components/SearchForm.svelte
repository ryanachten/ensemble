<script lang="ts">
	import { LayoutKeys } from '../graph/layout';

	export let className: string | undefined = undefined;
	export let layoutKey: LayoutKeys;
	export let bandName: string;
	export let degreesOfSeparation: number;
	export let onSubmitForm: () => void;
	export let onCenterGraph: () => void;
</script>

<form class={`${className}`} on:submit={onSubmitForm}>
	<div class="mb-4">
		<input class="input input-bordered mr-2" type="text" bind:value={bandName} />
		<input class="input input-bordered mr-2" type="number" bind:value={degreesOfSeparation} />
		<select class="select mr-2" name="layout options" id="layoutOptions" bind:value={layoutKey}>
			<option value={LayoutKeys.COSE}>Cose</option>
			<option value={LayoutKeys.BREADTH_FIRST}>Breadth-first</option>
		</select>
		<button class="btn btn-primary" type="submit">Search</button>
	</div>
	<div>
		<button
			class="btn"
			on:click={(e) => {
				e.preventDefault();
				onCenterGraph();
			}}>Center</button
		>
	</div>
</form>
