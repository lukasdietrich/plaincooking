import { defineConfig } from 'unocss';

export default defineConfig({
	rules: [[/^thickness-(\d+)$/, ([, d]) => ({ 'text-decoration-thickness': `${+d / 8}rem` })]]
});
