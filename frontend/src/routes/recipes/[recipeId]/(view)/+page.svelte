<script lang="ts">
	import type { PageData } from './$types';
	import { page } from '$app/stores';
	import { resolveRoute } from '$app/paths';
	import { t } from '$lib/i18n';
	import { Step } from '$lib/components/recipe';
	import { ActionPortal, Action } from '$lib/components/actions';

	export let data: PageData;
</script>

<ActionPortal>
	<Action
		href={resolveRoute('/recipes/[recipeId]/shopping', $page.params)}
		title={$t('actions.recipe.shopping')}
	>
		<i class="icon-shopping-cart"></i>
	</Action>
</ActionPortal>

{#each data.recipe.steps as tokens, index}
	<Step {index} {tokens} />
{/each}
