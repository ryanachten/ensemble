import type { Stylesheet } from 'cytoscape';
import bandStyles from './band.styles';
import commonStyles from './common.styles';
import genreStyles from './genre.styles';

const styles: Stylesheet[] = [...commonStyles, ...bandStyles, ...genreStyles];

export default styles;
