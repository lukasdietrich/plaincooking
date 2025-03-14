import type { components } from '$lib/api';
import { registerMiddleware } from '$lib/api';
import { notify } from './store';

export { notify };
export { default as Notifications } from './Notifications.svelte';

registerMiddleware({
	async onResponse({ response }) {
		if (response.status >= 300) {
			const error = <components['schemas']['ApiError']>await response.clone().json();

			notify({
				type: 'error',
				text: `api.error.${error.code}`
			});
		}

		return undefined;
	}
});
