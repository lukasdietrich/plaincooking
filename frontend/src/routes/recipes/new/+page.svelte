<script lang="ts">
	import { goto } from '$app/navigation';
	import { page } from '$app/stores';
	import { resolveRoute } from '$lib/routing';
	import { client } from '$lib';
	import { t } from '$lib/i18n';
	import { ActionPortal, Action } from '$lib/components/actions';
	import { Editor } from '$lib/components/editor';

	let content = '';

	async function save() {
		const { id } = await client.recipes.createRecipe(content);
		const url = resolveRoute('/recipes/[recipeId]', $page, { params: { recipeId: id } });

		await goto(url, { invalidateAll: true });
	}
</script>

<ActionPortal>
	<Action href="/" title={$t('actions.back')}>
		<i class="icon-undo-2"></i>
	</Action>

	<Action on:click={save} title={$t('actions.save')}>
		<i class="icon-save"></i>
	</Action>
</ActionPortal>

<Editor bind:value={content} />
