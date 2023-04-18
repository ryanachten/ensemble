import type { Stylesheet } from 'cytoscape';

const style: Stylesheet[] = [
	{
		selector: 'node[label]',
		style: {
			label: 'data(label)',
			'font-size': 10,
			'text-valign': 'center',
			'text-halign': 'center'
		}
	},

	{
		selector: 'edge[label]',
		style: {
			label: 'data(label)',
			width: 3,
			'font-size': 5,
			'edge-text-rotation': 'autorotate'
		} as any
	}
];

export default style;
