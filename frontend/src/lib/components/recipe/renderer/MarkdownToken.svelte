<script lang="ts">
	import type { Token } from 'marked';
	import Ingredient from './Ingredient.svelte';

	export let token: Token[] | Token = [];
</script>

{#if Array.isArray(token)}
	{#each token as token}
		<svelte:self {token} />
	{/each}
{:else if token.type === 'heading'}
	<h1 class="text-2xl font-semibold">
		<svelte:self token={token.tokens} />
	</h1>
{:else if token.type === 'paragraph'}
	<p class="py-2">
		<svelte:self token={token.tokens} />
	</p>
{:else if token.type === 'codespan'}
	<Ingredient text={token.text} />
{:else if token.type === 'text'}
	{token.text}
{/if}
