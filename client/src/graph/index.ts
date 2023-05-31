import cytoscape, { type ElementsDefinition, type LayoutOptions } from 'cytoscape';
import popper from 'cytoscape-popper';
import style from './styles';

cytoscape.use(popper);

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
