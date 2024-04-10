<script lang="ts">
	import type { components } from '$lib/api';
	import { page } from '$app/stores';
	import icon from '$lib/images/icon.png';
	import { Action } from '$lib/components/actions';
	import User from './User.svelte';

	export let user: components['schemas']['UserInfo'] | null = null;
</script>

<header class="sticky top-0 z-500 flex items-center h-12 sm:h-14 bg-emerald-600 text-white">
	<a class="flex items-center h-full px-3 sm:px-5 py-1 space-x-1 sm:space-x-3" href="/">
		<img class="h-full aspect-square" src={icon} alt="" />

		<h1 class="sm:text-xl font-semibold py-3">Plaincooking</h1>
	</a>

	<div id="actions" class="flex flex-1 justify-end h-full"></div>

	{#if user}
		<User {user} />
	{:else}
		<Action href="/oauth2/redirect?origin={$page.url.pathname}">
			<i class="icon-log-in"></i>
		</Action>
	{/if}
</header>
