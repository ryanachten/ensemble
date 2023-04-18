import cytoscape, {
	type EdgeDefinition,
	type NodeDataDefinition,
	type NodeDefinition
} from 'cytoscape';
import type { Graph } from '../models/api';
import layout from './layout';
import style from './style';

export default (graphData: Graph, container: HTMLElement) => {
	// TODO: pretty inefficient that we need to iterate through the entire graph to convert it into a compliant version
	// maybe we should just update the contract to return this from the backend
	const edges: EdgeDefinition[] = [];
	const nodes: NodeDefinition[] = Object.keys(graphData.vertices).map((vertexKey) => {
		const vertex = graphData.vertices[vertexKey as any];
		if (vertex.edges) {
			Object.keys(vertex.edges).map((edgeKey) => {
				const edge = vertex.edges![edgeKey as any];
				edges.push({
					data: {
						source: vertexKey,
						target: edgeKey,
						label: edge.label
					}
				});
			});
		}
		const nodeData: NodeDataDefinition = {
			id: vertexKey,
			label: vertexKey
		};
		return { data: nodeData };
	});
	return cytoscape({
		container,
		elements: {
			nodes,
			edges
		},
		style,
		layout
	});
};
