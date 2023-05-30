/** @type {import('tailwindcss').Config} */
export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	daisyui: {
		themes: [
			{
				ensemble: {
					'base-200': '#f5f5f5',
					primary: '#af91ff',
					secondary: '#50c1f1',
					accent: '#68e8c8',
					neutral: '#264653',
					'neutral-content': '#FFFFFF',
					'base-100': '#FFFFFF',
					info: '#50c1f1',
					success: '#68e8c8',
					warning: '#f5bd56',
					error: '#e76f51'
				}
			}
		]
	},
	theme: {
		extend: {
			fontFamily: {
				sans: ['Open Sans', 'ui-sans-serif']
			}
		}
	},
	plugins: [require('daisyui')]
};
