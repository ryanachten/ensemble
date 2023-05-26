import { writable } from 'svelte/store';
import type { NodeMetadata } from './interfaces';

export const nodePath = writable<NodeMetadata[]>([]);
export const confirmedNodePath = writable<NodeMetadata[]>([]);
