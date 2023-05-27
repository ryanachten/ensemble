import { writable } from 'svelte/store';
import type { NodeMetadata } from './interfaces';
import type { ElementsDefinition } from 'cytoscape';
import { LayoutKeys } from './graph/layout';
import type { Resource } from './api';

export const degreesOfSeparation = writable<number>(3);
export const layoutKey = writable<LayoutKeys>(LayoutKeys.COSE);
export const searchTerm = writable<string>('Black Flag');
export const resource = writable<Resource>('bands');

export const bands = writable<string[]>([]);
export const artists = writable<string[]>([]);
export const genres = writable<string[]>([]);
export const graphData = writable<ElementsDefinition>();

export const selectedItem = writable<string>();
export const nodePath = writable<NodeMetadata[]>([]);
export const confirmedNodePath = writable<NodeMetadata[]>([]);
