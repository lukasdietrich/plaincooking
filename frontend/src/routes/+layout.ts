import type { LayoutLoad } from './$types';
import { waitLocale } from '$lib/i18n';

export const prerender = true;
export const ssr = false;

export const load: LayoutLoad = async () => {
	await waitLocale();
};
