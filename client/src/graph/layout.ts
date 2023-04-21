import type { LayoutOptions, NodeSingular, Position } from 'cytoscape';

export enum LayoutKeys {
	COSE,
	BREADTH_FIRST
}

export const coseLayout: LayoutOptions = {
	name: 'cose',

	// Whether to animate while running the layout
	// true : Animate continuously as the layout is running
	// false : Just show the end result
	// 'end' : Animate with the end result, from the initial positions to the end positions
	animate: true,
	animationEasing: undefined, // Easing of the animation for animate:'end'
	animationDuration: undefined, // The duration of the animation for animate:'end'

	// A function that determines whether the node should be animated
	// All nodes animated by default on animate enabled
	// Non-animated nodes are positioned immediately when the layout starts
	animateFilter: () => true,

	// The layout animates only after this many milliseconds for animate:true
	// (prevents flashing on fast runs)
	animationThreshold: 250,
	refresh: 20, // Number of iterations between consecutive screen positions update
	fit: true, // Whether to fit the network view after when done
	padding: 300, // Padding on fit
	boundingBox: undefined, // Constrain layout bounds; { x1, y1, x2, y2 } or { x1, y1, w, h }
	nodeDimensionsIncludeLabels: false, // Excludes the label when calculating node bounding boxes for the layout algorithm
	randomize: false, // Randomize the initial positions of the nodes (true) or use existing positions (false)
	componentSpacing: 200, // Extra spacing between components in non-compound graphs
	nodeRepulsion: () => 2048, // Node repulsion (non overlapping) multiplier
	nodeOverlap: 10, // Node repulsion (overlapping) multiplier
	idealEdgeLength: () => 32, // Ideal edge (non nested) length
	edgeElasticity: () => 32, // Divisor to compute edge forces
	nestingFactor: 1.2, // Nesting factor (multiplier) to compute ideal edge length for nested edges
	gravity: 1, // Gravity force (constant)
	numIter: 1000, // Maximum number of iterations to perform
	initialTemp: 1000, // Initial temperature (maximum node displacement)
	coolingFactor: 0.99, // Cooling factor (how the temperature is reduced between consecutive iterations
	minTemp: 1.0 // Lower temperature threshold (below this point the layout will end)
};

export const breadthFirstLayout: LayoutOptions = {
	name: 'breadthfirst',
	fit: true, // whether to fit the viewport to the graph
	directed: false, // whether the tree is directed downwards (or edges can point in any direction if false)
	padding: 30, // padding on fit
	circle: false, // put depths in concentric circles if true, put depths top down if false
	grid: false, // whether to create an even grid into which the DAG is placed (circle:false only)
	spacingFactor: 1.75, // positive spacing factor, larger => more space between nodes (N.B. n/a if causes overlap)
	boundingBox: undefined, // constrain layout bounds; { x1, y1, x2, y2 } or { x1, y1, w, h }
	avoidOverlap: true, // prevents node overlap, may overflow boundingBox if not enough space
	nodeDimensionsIncludeLabels: false, // Excludes the label when calculating node bounding boxes for the layout algorithm
	roots: undefined, // the roots of the trees
	maximal: false, // whether to shift nodes down their natural BFS depths in order to avoid upwards edges (DAGS only)
	depthSort: undefined, // a sorting function to order nodes at equal depth. e.g. function(a, b){ return a.data('weight') - b.data('weight') }
	animate: false, // whether to transition the node positions
	animationDuration: 500, // duration of animation in ms if enabled
	animationEasing: undefined, // easing of animation if enabled,
	animateFilter: () => true, // a function that determines whether the node should be animated.  All nodes animated by default on animate enabled.  Non-animated nodes are positioned immediately when the layout starts
	ready: undefined, // callback on layoutready
	stop: undefined, // callback on layoutstop
	transform: (_node: NodeSingular, position: Position) => position // transform a given node position. Useful for changing flow direction in discrete layouts
};

export const layouts: Record<LayoutKeys, LayoutOptions> = {
	[LayoutKeys.COSE]: coseLayout,
	[LayoutKeys.BREADTH_FIRST]: breadthFirstLayout
};
