<script lang="ts">
	import type { Recipe } from '$lib/recipe';
	import { t } from '$lib/i18n';
	import TokenRenderer from './TokenRenderer.svelte';

	export let recipe: Recipe;

	$: metadata = recipe.metadata;
</script>

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
	</div>

	<TokenRenderer token={recipe.intro} />
</header>
