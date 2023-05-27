<script lang="ts">
	import type { Core, ElementsDefinition } from 'cytoscape';
	import type { Instance } from '@popperjs/core';

	import graph from '../graph';
	import { LayoutKeys, layouts } from '../graph/layout';
	import Tooltip from './Tooltip.svelte';
	import { confirmedNodePath, graphData, isLoading } from '../stores';

	export let layoutKey: LayoutKeys;
	export let className: string | undefined = undefined;
	export let selectedId: string | undefined = undefined;
	export const centerGraph = () => {
		cytoscape?.center().fit();
	};

	let container: HTMLDivElement;
	let cytoscape: Core | null = null;
	let popper: Instance | undefined;
	const CONTAINER_ID = 'tooltipWrapper';
	const PATH_NODE_CLASS = 'path-node';
	const NONE_PATH_NODE_CLASS = 'non-path-node';

	$: updateLayout(layoutKey);
	$: selectItem(selectedId);
	$: renderGraph($graphData);

	confirmedNodePath.subscribe((path) => {
		// If there's less than one node, reset styles
		if (path.length <= 1) {
			cytoscape?.elements().each((node) => {
				node.removeClass(PATH_NODE_CLASS);
				node.removeClass(NONE_PATH_NODE_CLASS);
			});
			return;
		}

		cytoscape?.elements().each((node) => {
			node.removeClass(PATH_NODE_CLASS);
			node.addClass(NONE_PATH_NODE_CLASS);
		});

		path.forEach((node, index) => {
			if (index >= path.length - 1) return;
			const target = cytoscape?.$id(path[index + 1].id);
			if (!target) return;

			cytoscape?.$id(node.id).addClass(PATH_NODE_CLASS);
			const floydWarshall = cytoscape?.elements().floydWarshall({
				weight: (node) => {
					// Prefer not to direct path relationships via genre nodes
					if (node.data('label') === 'genre') return 2;
					return 1;
				}
			});
			const pathToTarget = floydWarshall?.path(cytoscape?.$id(node.id)!, target);
			pathToTarget?.each((node) => {
				node.addClass(PATH_NODE_CLASS);
				node.removeClass(NONE_PATH_NODE_CLASS);
			});
		});
	});

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
		cytoscape?.on('tapunselect destroy', () => popper?.destroy());
	};
</script>

<div class={className} bind:this={container} />
