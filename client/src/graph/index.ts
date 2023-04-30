import cytoscape, { type ElementsDefinition, type LayoutOptions } from 'cytoscape';
import style from './styles';

export default ({
	data,
	container,
	layout
}: {
	data: ElementsDefinition;
	container: HTMLElement;
	layout: LayoutOptions;
}) => {
	return cytoscape({
		container,
		elements: data,
		style,
		layout
	});
};
