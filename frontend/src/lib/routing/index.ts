import type { Page } from '@sveltejs/kit';
import { resolveRoute as resolve } from '$app/paths';

interface Options {
	keepParams: boolean;
	keepQuery: boolean;

	params: Parameters<typeof resolve>[1];
}

const buildParams = (options: Options, $page: Page) => {
	if (options.keepParams) {
		return {
			...$page.params,
			...options.params
		};
	}

	return options.params;
};

const buildId = (options: Options, $page: Page, id: string) => {
	if (options.keepQuery) {
		const { search } = $page.url;
		return `${id}${search}`;
	}

	return id;
};

export const resolveRoute = (id: string, $page: Page, options?: Partial<Options>) => {
	const finalOptions = {
		keepParams: true,
		keepQuery: false,
		params: {},

		...options
	};

	const finalParams = buildParams(finalOptions, $page);
	const finalId = buildId(finalOptions, $page, id);

	return resolve(finalId, finalParams);
};
