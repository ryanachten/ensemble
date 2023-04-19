import type cytoscape from 'cytoscape';
import type { Stylesheet } from 'cytoscape';

const style: Stylesheet[] = [
	{
		selector: 'node',
		style: {
			'background-image': 'data(imageUrl)',
			'background-fit': 'cover',
			'border-width': 1
		}
	},
	{
		selector: 'node[label]',
		style: {
			label: 'data(label)',
			'font-size': 10,
			'text-valign': 'bottom',
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
		} as cytoscape.Css.Edge
	},
	{
		selector: '[type = "band"]',
		style: {
			backgroundColor: 'red',
			'border-color': 'red'
		}
	},
	{
		selector: '[type = "artist"]',
		style: {
			backgroundColor: 'blue',
			'border-color': 'blue'
		}
	}
];

export default style;
