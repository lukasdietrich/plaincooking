import type { LayoutLoad } from './$types';
import { waitLocale } from '$lib/i18n';
import { createApi } from '$lib/api';

export const prerender = true;
export const ssr = false;

export const load: LayoutLoad = async ({ fetch }) => {
	const [user] = await Promise.all([tryReadUserInfo(fetch), waitLocale()]);

	return {
		user
	};
};

async function tryReadUserInfo(fetch: typeof globalThis.fetch) {
	const { readUserInfo } = createApi(fetch, true);

	try {
		return await readUserInfo();
	} catch {
		return null;
	}
}
