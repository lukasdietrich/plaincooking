<script lang="ts">
	import type { Recipe, Section, Tokens, Token } from '$lib/recipe';
	import FractionalQuantity from './FractionalQuantity.svelte';

	export let recipe: Recipe;

	$: ingredients = findIngredients(recipe.steps);

	function findIngredients(steps: Section[]): Tokens.Ingredient[] {
		return steps
			.flat()
			.map(explodeToken)
			.flat()
			.filter(({ type }) => type === 'ingredient')
			.map((token) => <Tokens.Ingredient>token)
			.toSorted(compareIngredients);
	}

	function explodeToken(token: Token): Token[] {
		if ('children' in token) {
			return [token, ...token.children.map(explodeToken)].flat();
		}

		return [token];
	}

	function compareIngredients(a: Tokens.Ingredient, b: Tokens.Ingredient) {
		return a.ingredient.localeCompare(b.ingredient);
	}
</script>

<table class="font-serif text-lg">
	<tbody>
		{#each ingredients as token}
			<tr class="odd:bg-slate-50">
				<td colspan={token.quantity ? undefined : 2} class="px-3 py-2">
					{token.ingredient}
				</td>

				{#if token.quantity}
					<td class="text-right px-3">
						<FractionalQuantity quantity={token.quantity} unit={token.unit} />
					</td>
				{/if}
			</tr>
		{/each}
	</tbody>
</table>
