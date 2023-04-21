import type { NodeDataDefinition } from 'cytoscape';

export enum NodeType {
	ARTIST = 'artist',
	BAND = 'band'
}

export interface NodeData extends NodeDataDefinition {
	type: NodeType;
	imageUrl?: string;
}
