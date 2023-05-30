import type cytoscape from 'cytoscape';
import type { Stylesheet } from 'cytoscape';
import { theme } from '../../theme';

const commonStyles: Stylesheet[] = [
	{
		selector: 'node',
		style: {
			'background-image': 'data(imageUrl)',
			'background-fit': 'cover',
			'border-width': 1
		}
	},
	{
		selector: '[label]',
		style: {
			'font-family': 'Open Sans',
			color: theme.neutral
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
	}
];

export default commonStyles;
