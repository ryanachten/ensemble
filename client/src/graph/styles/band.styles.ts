import type { Stylesheet } from 'cytoscape';
import { theme } from '../../theme';

const bandStyles: Stylesheet[] = [
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

export default bandStyles;
