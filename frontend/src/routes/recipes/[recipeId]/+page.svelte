<script lang="ts">
	import type { PageData } from './$types';
	import { page } from '$app/stores';
	import { resolveRoute } from '$app/paths';
	import { t } from '$lib/i18n';
	import { ActionPortal, Action } from '$lib/components/actions';
	import { Recipe } from '$lib/components/recipe';

	export let data: PageData;
</script>

<ActionPortal>
	<Action href={resolveRoute('/', $page.params)} title={$t('actions.back')}>
		<i class="icon-undo-2"></i>
	</Action>

	<Action href={resolveRoute('/recipes/[recipeId]/edit', $page.params)} title={$t('actions.edit')}>
		<i class="icon-file-pen-line"></i>
	</Action>
</ActionPortal>

<Recipe recipe={data.recipe} />
