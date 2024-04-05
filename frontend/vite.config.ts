import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import unocss from '@unocss/svelte-scoped/vite';

export default defineConfig({
	plugins: [unocss(), sveltekit()],

	build: {
		// The only chunk larger than 500kB contains the codemirror editor.
		// It is already code-split properly and only loaded when navigating to the edit page.
		chunkSizeWarningLimit: 768
	},

	server: {
		proxy: {
			'/api': 'http://localhost:8080',
			'/oauth2': 'http://localhost:8080'
		}
	}
});
