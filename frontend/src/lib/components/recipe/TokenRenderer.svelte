<script lang="ts">
	import type { Token } from '$lib/recipe';
	import Ingredient from './Ingredient.svelte';

	export let token: Token[] | Token = [];
</script>

{#if Array.isArray(token)}
	{#each token as token, index (index)}
		<svelte:self {token} />
	{/each}
{:else if token.type === 'em'}
	<em>
		<svelte:self token={token.children} />
	</em>
{:else if token.type === 'heading'}
	<h1 class="text-2xl font-semibold">
		<svelte:self token={token.children} />
	</h1>
{:else if token.type === 'ingredient'}
	<Ingredient {token} />
{:else if token.type === 'paragraph'}
	<p class="py-2 font-serif text-lg">
		<svelte:self token={token.children} />
	</p>
{:else if token.type === 'strong'}
	<strong>
		<svelte:self token={token.children} />
	</strong>
{:else if token.type === 'text'}
	<!-- eslint-disable-next-line svelte/no-at-html-tags -->
	{@html token.text}
{/if}
