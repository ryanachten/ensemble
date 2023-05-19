<script lang="ts">
	import type { Core, ElementsDefinition } from 'cytoscape';
	import type { Instance } from '@popperjs/core';

	import graph from '../graph';
	import { LayoutKeys, layouts } from '../graph/layout';
	import Tooltip from './Tooltip.svelte';

	export let layoutKey: LayoutKeys;
	export let data: ElementsDefinition | null;
	export let className: string | undefined = undefined;
	export let selectedId: string | undefined = undefined;
	export const centerGraph = () => {
		cytoscape?.center().fit();
	};

	let container: HTMLDivElement;
	let cytoscape: Core | null = null;
	let popper: Instance | undefined;
	const CONTAINER_ID = 'tooltipWrapper';

	$: updateLayout(layoutKey);
	$: selectItem(selectedId);
	$: renderGraph(data);

	const renderGraph = (updatedData: ElementsDefinition | null) => {
		if (!updatedData) return;

		cytoscape?.destroy();
		cytoscape = graph({ data: updatedData, container, layout: layouts[layoutKey] });
		cytoscape?.on('layoutstop', () => centerGraph());
		cytoscape?.on('select', 'node', (e) => selectItem(e.target.id()));
	};

	const updateLayout = (newLayoutKey: LayoutKeys) => {
		const newLayout = layouts[newLayoutKey];
		const updatedLayout = cytoscape?.layout(newLayout);
		updatedLayout?.run();
	};

	const deletePopover = () => {
		popper?.destroy();
		const remainingContainer = document.getElementById(CONTAINER_ID);
		remainingContainer && document.body.removeChild(remainingContainer);
	};

	const selectItem = (id: string | undefined) => {
		if (!id) return;
		deletePopover();

		const selectedNode = cytoscape?.$id(id);
		const selection = selectedNode?.connectedEdges().connectedNodes();
		cytoscape?.fit(selection);

		popper = selectedNode?.popper({
			content: () => {
				const toolTipWrapper = document.createElement('div');
				toolTipWrapper.id = CONTAINER_ID;
				new Tooltip({
					target: toolTipWrapper,
					props: {
						data: selectedNode.data()
					}
				});

				document.body.appendChild(toolTipWrapper);

				return toolTipWrapper;
			}
		});
		let update = () => {
			popper?.update();
		};

		selectedNode?.on('position', update);
		cytoscape?.on('pan zoom resize', update);
	};
</script>

<div class={className} bind:this={container} />
