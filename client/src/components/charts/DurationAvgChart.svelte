<script lang="ts">
	import { Chart } from 'chart.js/auto';

	import { generateDataset, type ChartParams } from '.';

	export let className: string;

	export const init = ({ results, labels, formattedLabels }: ChartParams) => {
		const durationAvg = generateDataset(results, labels, {
			columnKey: 'durationAvg',
			formatValue: (value) => {
				const parsedVal = parseFloat(value);
				if (value.includes('ms')) return parsedVal;
				if (value.includes('s')) return parsedVal * 1000;
				if (value.includes('m') || value.includes('min')) return parsedVal * 60000;
				return parsedVal;
			}
		});
		new Chart('durationAvg', {
			type: 'bar',
			data: {
				labels: formattedLabels,
				datasets: durationAvg
			},
			options: {
				responsive: true,
				plugins: {
					title: {
						display: true,
						text: 'Average duration'
					}
				}
			}
		});
	};
</script>

<div class="relative {className}">
	<canvas id="durationAvg" />
</div>
