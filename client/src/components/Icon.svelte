<script lang="ts">
	import feather, { type FeatherIconNames } from 'feather-icons';
	export let name: FeatherIconNames;
	export let className: string = '';

	$: icon = feather.icons[name];
	$: if (icon) {
		icon.attrs['height'] = '1em';
		icon.attrs['width'] = '1em';
		icon.attrs['class'] = className;
	}
</script>

{@html icon.toSvg()}
