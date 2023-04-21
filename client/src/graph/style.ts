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
			width: 2,
			'font-size': 5,
			'edge-text-rotation': 'autorotate'
		} as cytoscape.Css.Edge
	},
	{
		selector: '[type = "band"]',
		style: {
			backgroundColor: 'purple',
			'border-color': 'purple'
		}
	},
	{
		selector: '[type = "artist"]',
		style: {
			backgroundColor: 'blue',
			'border-color': 'blue'
		}
	},
	{
		selector: '[label = "member"], [label = "member of"]',
		style: {
			'line-color': 'green'
		}
	},
	{
		selector: '[label = "past member"], [label = "formerly of"]',
		style: {
			'line-color': 'red',
			'line-style': 'dashed',
			'line-cap': 'round',
			'line-dash-pattern': [1, 5]
		}
	}
];

export default style;
