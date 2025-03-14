import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: vitePreprocess(),

	kit: {
		adapter: adapter(),

		paths: {
			relative: false
		},

		prerender: {
			entries: [
				'*',
				'/recipes/[recipeId]',
				'/recipes/[recipeId]/shopping',
				'/recipes/[recipeId]/edit'
			]
		}
	}
};

export default config;
