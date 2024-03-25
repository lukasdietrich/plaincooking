<script lang="ts">
	import type { Recipe } from '$lib/recipe';
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { t } from '$lib/i18n';
	import TokenRenderer from './TokenRenderer.svelte';

	export let recipe: Recipe;

	$: metadata = recipe.metadata;
	$: servings = metadata.servings;

	function updateServings(delta: number) {
		const url = new URL($page.url);
		url.searchParams.set('servings', String(servings + delta));
		goto(url);
	}
</script>

<header>
	<div class="flex flex-wrap mb-5 text-sm badges">
		<div class="flex items-center space-x-2 rounded-full bg-yellow-200 text-yellow-900 px-3 py-1">
			<i class="icon-utensils"></i>

			{#if servings > 1}
				<button on:click={() => updateServings(-1)}>
					<i class="icon-minus"></i>
				</button>
			{/if}

			<span>
				{$t('recipe.servings', { values: { n: metadata.servings } })}
			</span>

			<button on:click={() => updateServings(1)}>
				<i class="icon-plus"></i>
			</button>
		</div>

		{#if metadata.source}
			<div class="flex items-center space-x-2 rounded-full bg-purple-200 text-purple-900 px-3 py-1">
				<i class="icon-earth"></i>
				<a href={metadata.source} target="_blank">{new URL(metadata.source).hostname}</a>
			</div>
		{/if}

		{#each metadata.tags as tag}
			<div class="flex items-center space-x-1 rounded-full bg-blue-200 text-blue-900 px-3 py-1">
				<i class="icon-hash"></i>
				<span>{tag}</span>
			</div>
		{/each}
	</div>

	<TokenRenderer token={recipe.intro} />
</header>

<style>
	.badges > div {
		--at-apply: mb-2 mr-2;
	}
</style>
