import type { ElementsDefinition, NodeDataDefinition } from 'cytoscape';

export enum NodeType {
	ARTIST = 'artist',
	BAND = 'band',
	GENRE = 'genre'
}

export interface NodeData extends NodeDataDefinition {
	type: NodeType;
	imageUrl?: string;
}

export const requestBandGraph = async (bandName: string, degreesOfSeparation: number) => {
	const res = await fetch(
		`http://localhost:8080/bands?name=${bandName}&degreesOfSeparation=${degreesOfSeparation}`
	);
	const data = (await res.json()) as ElementsDefinition;

	const bands: string[] = [];
	const artists: string[] = [];
	const genres: string[] = [];

	data.nodes.forEach((node) => {
		const data = node.data as NodeData;
		if (data.type === NodeType.ARTIST && data.id) {
			bands.push(data.id);
		}
		if (data.type === NodeType.BAND && data.id) {
			artists.push(data.id);
		}
		if (data.type === NodeType.GENRE && data.id) {
			genres.push(data.id);
		}
	});

	artists.sort();
	bands.sort();
	genres.sort();

	return {
		data,
		bands,
		artists,
		genres
	};
};
