import { defineConfig } from 'unocss';

export default defineConfig({
	theme: {
		fontFamily: {
			sans: `'Readex Pro Variable', sans-serif`,
			serif: `'Crimson Pro Variable', serif`,
			mono: `'Inconsolata Variable', monospace`
		}
	},

	rules: [[/^thickness-(\d+)$/, ([, d]) => ({ 'text-decoration-thickness': `${+d / 8}rem` })]]
});
