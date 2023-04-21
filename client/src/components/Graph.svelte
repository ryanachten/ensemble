<script lang="ts">
	import graph from '../graph';
	import type { Core, ElementsDefinition } from 'cytoscape';
	import { LayoutKeys, layouts } from '../graph/layout';

	export let layoutKey: LayoutKeys;
	export let data: ElementsDefinition | null;
	export let className: string | undefined = undefined;
	export let selectedId: string | undefined = undefined;

	let container: HTMLDivElement;
	let cytoscape: Core | null = null;

	$: updateLayout(layoutKey);
	$: selectItem(selectedId);
	$: renderGraph(data);

	const renderGraph = (updatedData: ElementsDefinition | null) => {
		if (!updatedData) return;

		cytoscape?.destroy();
		cytoscape = graph({ data: updatedData, container, layout: layouts[layoutKey] });
	};

	const updateLayout = (newLayoutKey: LayoutKeys) => {
		const newLayout = layouts[newLayoutKey];
		const updatedLayout = cytoscape?.layout(newLayout);
		updatedLayout?.run();
	};

	const selectItem = (id: string | undefined) => {
		if (!id) return;

		const selection = cytoscape?.$id(id).connectedEdges().connectedNodes();
		cytoscape?.fit(selection);
		selectedId = undefined;
	};
</script>

<div class={className} bind:this={container} />
