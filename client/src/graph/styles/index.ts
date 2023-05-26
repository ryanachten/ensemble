import type { Stylesheet } from 'cytoscape';
import bandStyles from './band.styles';
import commonStyles from './common.styles';
import genreStyles from './genre.styles';

const styles: Stylesheet[] = [
	...commonStyles,
	...bandStyles,
	...genreStyles,
	{
		selector: '.node-path',
		style: {
			backgroundColor: 'red',
			'border-color': 'red',
			'line-color': 'red'
		}
	}
];

export default styles;
