<script lang="ts">
	import type { NodeType } from '../interfaces';

	export let data: { id: string; label: string; type: NodeType };
</script>

<div class="card bg-base-100 shadow-xl">
	<div class="card-body">
		<h2 class="card-title">{data.label}</h2>
		<span class="badge">{data.type}</span>
	</div>
</div>
