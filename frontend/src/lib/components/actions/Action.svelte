<script lang="ts">
	import { user } from '$lib/auth';

	export let href: string | null = null;
	export let title: string | null = null;
	export let requireAuth: boolean = false;

	$: show = !requireAuth || Boolean($user);
</script>

{#if show}
	{#if href}
		<a {href} {title} on:click>
			<slot />
		</a>
	{:else}
		<button {title} on:click>
			<slot />
		</button>
	{/if}
{/if}

<style lang="css">
	a,
	button {
		--at-apply: 'flex h-full px-4 sm:px-5 items-center bg-emerald-700 text-xl ring-emerald-900 transition hover:(bg-emerald-100 text-emerald-900 ring-2 z-2)'; /**/
	}
</style>
