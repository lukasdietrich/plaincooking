import type { components } from '$lib/api';
import { derived } from 'svelte/store';
import { page } from '$app/stores';

export const user = derived(page, ($page) => {
	return <components['schemas']['UserInfo'] | null>$page.data.user;
});
