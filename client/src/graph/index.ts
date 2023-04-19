import cytoscape, { type ElementsDefinition } from 'cytoscape';
import layout from './layout';
import style from './style';

export default (graphData: ElementsDefinition, container: HTMLElement) => {
	return cytoscape({
		container,
		elements: graphData,
		style,
		layout
	});
};
