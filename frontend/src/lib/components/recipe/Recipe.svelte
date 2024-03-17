<script lang="ts">
	import type { Recipe } from '$lib/recipe';
	import { t } from '$lib/i18n';
	import TokenRenderer from './TokenRenderer.svelte';
	import Step from './Step.svelte';
	import ShoppingList from './ShoppingList.svelte';

	enum Mode {
		Steps,
		ShoppingList
	}

	export let recipe: Recipe;

	let mode: Mode = Mode.Steps;

	$: metadata = recipe.metadata;

	function toggleMode() {
		mode = mode === Mode.Steps ? Mode.ShoppingList : Mode.Steps;
	}
</script>

<article class="flex flex-col space-y-5">
	<header>
		<div class="flex space-x-2 mb-5 text-sm">
			<div class="flex space-x-2 rounded-full bg-yellow-200 text-yellow-900 px-3 py-1">
				<i class="icon-utensils"></i>
				<span>
					{$t('recipe.servings', { values: { n: metadata.servings } })}
				</span>
			</div>

			{#each metadata.tags as tag}
				<div class="flex space-x-1 rounded-full bg-blue-200 text-blue-900 px-3 py-1">
					<i class="icon-hash"></i>
					<span>{tag}</span>
				</div>
			{/each}

			<div class="!ml-auto">
				<button class="bg-gray-200 rounded-full px-3 py-1" on:click={toggleMode}>
					<i class="icon-shopping-cart"></i>
				</button>
			</div>
		</div>

		<TokenRenderer token={recipe.intro} />
	</header>

	{#if mode === Mode.Steps}
		{#each recipe.steps as tokens, index}
			<Step {index} {tokens} />
		{/each}
	{:else}
		<ShoppingList {recipe} />
	{/if}
</article>
