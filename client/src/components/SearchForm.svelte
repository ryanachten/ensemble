<script lang="ts">
	import { LayoutKeys } from '../graph/layout';
	import { degreesOfSeparation, resource, layoutKey, searchTerm } from '../stores';

	export let className: string | undefined = undefined;

	export let onSubmitForm: () => void;
	export let onCenterGraph: () => void;
</script>

<form class={`${className}`} on:submit={onSubmitForm}>
	<div class="grid grid-cols-2 gap-4">
		<div class="form-control">
			<label for="search-term" class="label">
				<span class="label-text">Search term</span>
			</label>
			<input
				id="search-term"
				class="input input-bordered input-sm"
				type="text"
				bind:value={$searchTerm}
			/>
		</div>
		<div class="form-control">
			<label for="resource" class="label">
				<span class="label-text">Search type</span>
			</label>
			<select class="select select-sm" id="resource" bind:value={$resource}>
				<option value="bands">Band</option>
				<option value="artists">Artist</option>
				<option value="genres">Genre</option>
			</select>
		</div>
		<div class="form-control">
			<label for="degreesOfSeparation" class="label">
				<span class="label-text">Degrees of separation</span>
			</label>
			<input
				class="input input-bordered input-sm"
				id="degreesOfSeparation"
				type="number"
				bind:value={$degreesOfSeparation}
			/>
		</div>
		<div class="form-control">
			<label for="layoutOptions" class="label">
				<span class="label-text">Layout mode</span>
			</label>
			<select class="select select-sm" id="layoutOptions" bind:value={$layoutKey}>
				<option value={LayoutKeys.COSE}>Cose</option>
				<option value={LayoutKeys.BREADTH_FIRST}>Breadth-first</option>
			</select>
		</div>
		<button class="btn btn-primary btn-sm" type="submit">Search</button>
		<button
			class="btn btn-sm"
			on:click={(e) => {
				e.preventDefault();
				onCenterGraph();
			}}>Center</button
		>
	</div>
</form>
