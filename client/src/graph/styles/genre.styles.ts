import type { Stylesheet } from 'cytoscape';
import { theme } from '../../theme';

const genreStyles: Stylesheet[] = [
	{
		selector: '[label = "stylistic origin"]',
		style: {
			'line-color': theme.accent,
			'line-style': 'dotted'
		}
	},
	{
		selector: '[label = "derivative form"]',
		style: {
			'line-color': theme.primary
		}
	},
	{
		selector: '[label = "subgenre"]',
		style: {
			'line-color': theme.secondary
		}
	},
	{
		selector: '[label = "fusion genre"]',
		style: {
			'line-color': theme.warning,
			'line-style': 'dashed',
			'line-cap': 'round',
			'line-dash-pattern': [1, 5]
		}
	}
];

export default genreStyles;
