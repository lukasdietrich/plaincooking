<script lang="ts">
	export let href: string | null = null;
</script>

{#if href}
	<a {href} on:click>
		<slot />
	</a>
{:else}
	<button on:click>
		<slot />
	</button>
{/if}

<style>
	a,
	button {
		--at-apply: 'flex h-full px-5 items-center ring-emerald-900 transition hover:(bg-emerald-100 text-emerald-900 ring-2)';
	}
</style>
