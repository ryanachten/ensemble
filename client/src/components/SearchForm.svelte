<script lang="ts">
	import type { Resource } from '../api';
	import { LayoutKeys } from '../graph/layout';

	export let className: string | undefined = undefined;
	export let layoutKey: LayoutKeys;
	export let name: string;
	export let resource: Resource;
	export let degreesOfSeparation: number;
	export let onSubmitForm: () => void;
	export let onCenterGraph: () => void;
</script>

<form class={`${className}`} on:submit={onSubmitForm}>
	<div class="flex items-end mb-4">
		<div class="form-control">
			<label for="search-term" class="label">
				<span class="label-text">Search term</span>
			</label>
			<input id="search-term" class="input input-bordered mr-2" type="text" bind:value={name} />
		</div>
		<div class="form-control">
			<label for="resource" class="label">
				<span class="label-text">Search type</span>
			</label>
			<select class="select mr-2" id="resource" bind:value={resource}>
				<option value="bands">Band</option>
				<option value="genres">Genre</option>
			</select>
		</div>
		<div class="form-control">
			<label for="degreesOfSeparation" class="label">
				<span class="label-text">Degrees of separation</span>
			</label>
			<input
				class="input input-bordered mr-2"
				id="degreesOfSeparation"
				type="number"
				bind:value={degreesOfSeparation}
			/>
		</div>
		<div class="form-control">
			<label for="layoutOptions" class="label">
				<span class="label-text">Layout mode</span>
			</label>
			<select class="select mr-2" id="layoutOptions" bind:value={layoutKey}>
				<option value={LayoutKeys.COSE}>Cose</option>
				<option value={LayoutKeys.BREADTH_FIRST}>Breadth-first</option>
			</select>
		</div>
		<button class="btn btn-primary mr-2" type="submit">Search</button>
		<button
			class="btn"
			on:click={(e) => {
				e.preventDefault();
				onCenterGraph();
			}}>Center</button
		>
	</div>
</form>
