import { NodeType } from './interfaces';

export const theme = {
	primary: '#af91ff',
	secondary: '#50c1f1',
	accent: '#68e8c8',
	neutral: '#264653',
	'base-100': '#FFFFFF',
	info: '#50c1f1',
	success: '#68e8c8',
	warning: '#f5bd56',
	error: '#e76f51'
};
type Theme = typeof theme;

export const nodeColourMap: Record<NodeType, keyof Theme> = {
	[NodeType.BAND]: 'primary',
	[NodeType.ARTIST]: 'secondary',
	[NodeType.GENRE]: 'accent'
};
