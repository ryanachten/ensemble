import type { NodeDataDefinition } from 'cytoscape';

export enum NodeType {
	ARTIST = 'artist',
	BAND = 'band',
	GENRE = 'genre'
}

export interface NodeData extends NodeDataDefinition {
	type: NodeType;
	label: string;
	imageUrl?: string;
}

export interface NodeMetadata {
	id: string;
	label: string;
	type: NodeType;
}
