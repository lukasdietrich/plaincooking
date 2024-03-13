<script lang="ts">
	import type { Recipe } from '$lib/recipe';
	import TokenRenderer from './TokenRenderer.svelte';
	import Step from './Step.svelte';

	export let recipe: Recipe;
</script>

<article class="flex flex-col space-y-5">
	<header>
		<TokenRenderer token={recipe.intro} />
	</header>

	{#each recipe.steps as tokens, index}
		<Step {index} {tokens} />
	{/each}
</article>
