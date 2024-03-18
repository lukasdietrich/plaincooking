<script lang="ts">
	import type { PageData } from './$types';
	import { page } from '$app/stores';
	import { resolveRoute } from '$app/paths';
	import { t } from '$lib/i18n';
	import { ShoppingList } from '$lib/components/recipe';
	import { ActionPortal, Action } from '$lib/components/actions';

	export let data: PageData;

	$: recipe = data.recipe;
</script>

<ActionPortal>
	<Action
		href={resolveRoute('/recipes/[recipeId]', $page.params)}
		title={$t('actions.recipe.steps')}
	>
		<i class="icon-list-todo"></i>
	</Action>
</ActionPortal>

<ShoppingList {recipe} />
