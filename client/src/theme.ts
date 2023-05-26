import { NodeType } from './interfaces';

export const theme = {
	primary: '#570DF8',
	secondary: '#F000B8',
	accent: '#37CDBE',
	neutral: '#3D4451',
	'base-100': '#FFFFFF',
	info: '#3ABFF8',
	success: '#36D399',
	warning: '#FBBD23',
	error: '#F87272'
};
type Theme = typeof theme;

export const nodeColourMap: Record<NodeType, keyof Theme> = {
	[NodeType.BAND]: 'primary',
	[NodeType.ARTIST]: 'secondary',
	[NodeType.GENRE]: 'accent'
};
