<script lang="ts">
	import { Chart } from 'chart.js/auto';

	import { generateDataset, type ChartParams } from '.';

	export let className: string;

	export const init = ({ results, labels, formattedLabels }: ChartParams) => {
		const checksPassed = generateDataset(results, labels, {
			columnKey: 'checksPassed',
			formatValue: (value) => parseFloat(value.replace('%', ''))
		});
		new Chart('checksPassed', {
			type: 'bar',
			data: {
				labels: formattedLabels,
				datasets: checksPassed
			},
			options: {
				responsive: true,
				plugins: {
					title: {
						display: true,
						text: 'Reliability'
					}
				}
			}
		});
	};
</script>

<div class="relative {className}">
	<canvas id="checksPassed" />
</div>
