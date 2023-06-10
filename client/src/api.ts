import { PUBLIC_API_URL } from '$env/static/public';
import type { ElementsDefinition } from 'cytoscape';
import { NodeType, type NodeData } from './interfaces';
import {
	bands,
	artists,
	genres,
	graphData,
	nodePath,
	confirmedNodePath,
	isLoading,
	hasError
} from './stores';

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
	isLoading.set(true);
	hasError.set(false);

	let data: ElementsDefinition;
	try {
		const res = await fetch(
			`${PUBLIC_API_URL}/${resource}?name=${name}&degreesOfSeparation=${degreesOfSeparation}`
		);
		if (!res.ok) {
			hasError.set(true);
			isLoading.set(false);
		}
		data = (await res.json()) as ElementsDefinition;
	} catch (error) {
		hasError.set(true);
		isLoading.set(false);
		return;
	}

	const _bands: NodeData[] = [];
	const _artists: NodeData[] = [];
	const _genres: NodeData[] = [];

	data.nodes.forEach((node) => {
		const data = node.data as NodeData;
		if (data.type === NodeType.ARTIST && data.id) {
			_artists.push(data);
		}
		if (data.type === NodeType.BAND && data.id) {
			_bands.push(data);
		}
		if (data.type === NodeType.GENRE && data.id) {
			_genres.push(data);
		}
	});

	const sortNodes = (a: NodeData, b: NodeData) => (a.label > b.label ? 1 : -1);
	_artists.sort(sortNodes);
	_bands.sort(sortNodes);
	_genres.sort(sortNodes);

	bands.set(_bands);
	artists.set(_artists);
	genres.set(_genres);
	graphData.set(data);
	nodePath.set([]);
	confirmedNodePath.set([]);
	isLoading.set(false);
};
