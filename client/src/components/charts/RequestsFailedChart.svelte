<script lang="ts">
	import { Chart } from 'chart.js/auto';

	import { generateDataset, type ChartParams } from '.';

	export let className: string;

	export const init = ({ results, labels, formattedLabels }: ChartParams) => {
		const requestsFailed = generateDataset(results, labels, {
			columnKey: 'requestsFailed',
			formatValue: (value) => parseFloat(value.replace('%', ''))
		});
		new Chart('requestsFailed', {
			type: 'bar',
			data: {
				labels: formattedLabels,
				datasets: requestsFailed
			},
			options: {
				responsive: true,
				plugins: {
					title: {
						display: true,
						text: 'Requests failed'
					}
				}
			}
		});
	};
</script>

<div class="relative {className}">
	<canvas id="requestsFailed" />
</div>
