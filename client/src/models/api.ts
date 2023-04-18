export enum VertexType {
	Band = 'Band',
	Artist = 'Artist'
}

export interface VertexData {
	type: VertexType;
	url?: string;
}

export interface Edge {
	label: string;
}

export interface Vertex {
	data: VertexData;
	edges?: Edge[];
}

export interface Graph {
	vertices: Vertex[];
}
