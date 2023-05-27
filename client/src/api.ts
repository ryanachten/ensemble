import { PUBLIC_API_URL } from '$env/static/public';
import type { ElementsDefinition } from 'cytoscape';
import { NodeType, type NodeData } from './interfaces';
import { bands, artists, genres, graphData, nodePath, confirmedNodePath } from './stores';

export type Resource = 'bands' | 'genres';
export type RequestGraphOptions = {
	resource: Resource;
	name: string;
	degreesOfSeparation: number;
};

export const requestGraph = async ({
	resource,
	name,
	degreesOfSeparation
}: RequestGraphOptions) => {
	const res = await fetch(
		`${PUBLIC_API_URL}/${resource}?name=${name}&degreesOfSeparation=${degreesOfSeparation}`
	);
	const data = (await res.json()) as ElementsDefinition;

	const _bands: string[] = [];
	const _artists: string[] = [];
	const _genres: string[] = [];

	data.nodes.forEach((node) => {
		const data = node.data as NodeData;
		if (data.type === NodeType.ARTIST && data.id) {
			_artists.push(data.id);
		}
		if (data.type === NodeType.BAND && data.id) {
			_bands.push(data.id);
		}
		if (data.type === NodeType.GENRE && data.id) {
			_genres.push(data.id);
		}
	});

	_artists.sort();
	_bands.sort();
	_genres.sort();

	bands.set(_bands);
	artists.set(_artists);
	genres.set(_genres);
	graphData.set(data);
	nodePath.set([]);
	confirmedNodePath.set([]);
};
