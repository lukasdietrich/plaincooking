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
		--at-apply: 'block px-5 py-3 transition-colors hover:bg-emerald-700';
	}
</style>
