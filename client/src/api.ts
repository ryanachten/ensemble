import { PUBLIC_API_URL } from '$env/static/public';
import type { ElementsDefinition } from 'cytoscape';
import { NodeType, type NodeData } from './interfaces';

export type Resource = 'bands' | 'genres';

export const requestGraph = async ({
	resource,
	name,
	degreesOfSeparation
}: {
	resource: Resource;
	name: string;
	degreesOfSeparation: number;
}) => {
	const res = await fetch(
		`${PUBLIC_API_URL}/${resource}?name=${name}&degreesOfSeparation=${degreesOfSeparation}`
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
