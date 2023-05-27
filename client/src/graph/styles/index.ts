import type { Stylesheet } from 'cytoscape';
import bandStyles from './band.styles';
import commonStyles from './common.styles';
import genreStyles from './genre.styles';
import { theme } from '../../theme';

const styles: Stylesheet[] = [
	...commonStyles,
	...bandStyles,
	...genreStyles,
	{
		selector: '.path-node',
		style: {
			backgroundColor: theme.warning,
			'border-color': theme.warning,
			'line-color': theme.warning,
			opacity: 1
		}
	},
	{
		selector: '.non-path-node',
		style: {
			opacity: 0.2
		}
	}
];

export default styles;
