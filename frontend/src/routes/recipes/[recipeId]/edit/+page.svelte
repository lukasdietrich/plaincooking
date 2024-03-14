<script lang="ts">
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';
	import { resolveRoute } from '$app/paths';
	import { page } from '$app/stores';
	import { client } from '$lib';
	import { t } from '$lib/i18n';
	import { ActionPortal, Action } from '$lib/components/actions';
	import { Editor } from '$lib/components/editor';

	export let data: PageData;
	let content = data.content;

	async function save() {
		const recipeId = $page.params.recipeId;

		await client.recipes.updateRecipe(recipeId, content);
		await goto(resolveRoute('/recipes/[recipeId]', $page.params), { invalidateAll: true });
	}
</script>

<ActionPortal>
	<Action href={resolveRoute('/recipes/[recipeId]', $page.params)} title={$t('actions.back')}>
		<i class="icon-undo-2"></i>
	</Action>

	<Action on:click={save} title={$t('actions.save')}>
		<i class="icon-save"></i>
	</Action>
</ActionPortal>

<div class="bg-gray-50 font-mono text-lg">
	<Editor bind:value={content} />
</div>
