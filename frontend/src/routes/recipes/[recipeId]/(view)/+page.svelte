<script lang="ts">
	import type { PageData } from './$types';
	import { page } from '$app/stores';
	import { resolveRoute } from '$lib/routing';
	import { t } from '$lib/i18n';
	import { Step } from '$lib/components/recipe';
	import { ActionPortal, Action } from '$lib/components/actions';

	export let data: PageData;

	$: shoppingHref = resolveRoute('/recipes/[recipeId]/shopping', $page, { keepQuery: true });
</script>

<ActionPortal>
	<Action href={shoppingHref} title={$t('actions.recipe.shopping')}>
		<i class="icon-shopping-cart"></i>
	</Action>
</ActionPortal>

{#each data.recipe.steps as tokens, index (index)}
	<Step {index} {tokens} />
{/each}
