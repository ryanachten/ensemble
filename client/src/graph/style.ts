import type cytoscape from 'cytoscape';
import type { Stylesheet } from 'cytoscape';
import { theme } from '../theme';

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
			backgroundColor: theme.primary,
			'border-color': theme.primary
		}
	},
	{
		selector: '[type = "artist"]',
		style: {
			backgroundColor: theme.secondary,
			'border-color': theme.secondary
		}
	},
	{
		selector: '[type = "genre"]',
		style: {
			backgroundColor: theme.accent,
			'border-color': theme.accent
		}
	},
	{
		selector: '[label = "genre"]',
		style: {
			'line-color': theme.accent,
			'line-style': 'dotted'
		}
	},
	{
		selector: '[label = "member"], [label = "member of"]',
		style: {
			'line-color': theme.secondary
		}
	},
	{
		selector: '[label = "past member"], [label = "formerly of"]',
		style: {
			'line-color': theme.secondary,
			'line-style': 'dashed',
			'line-cap': 'round',
			'line-dash-pattern': [1, 5]
		}
	}
];

export default style;
