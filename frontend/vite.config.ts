import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import unocss from '@unocss/svelte-scoped/vite';

export default defineConfig({
	plugins: [unocss(), sveltekit()],
	server: {
		proxy: {
			'/api': 'http://localhost:8080'
		}
	}
});
