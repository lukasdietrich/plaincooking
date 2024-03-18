<script lang="ts">
	import type { PageData } from './$types';
	import { goto } from '$app/navigation';
	import { resolveRoute } from '$app/paths';
	import { page } from '$app/stores';
	import { client } from '$lib';
	import { t } from '$lib/i18n';
	import { ActionPortal, Action } from '$lib/components/actions';
	import { Editor } from '$lib/components/editor';
	import Modal from '$lib/components/Modal.svelte';

	const BooleanModal = Modal<boolean>;

	export let data: PageData;

	let content = data.content;
	let deleteModal: Modal<boolean>;

	async function deleteRecipe() {
		if (await deleteModal.show()) {
			const recipeId = $page.params.recipeId;

			await client.recipes.deleteRecipe(recipeId);
			await goto('/', { invalidateAll: true });
		}
	}

	async function saveRecipe() {
		const recipeId = $page.params.recipeId;

		await client.recipes.updateRecipe(recipeId, content);
		await goto(resolveRoute('/recipes/[recipeId]', $page.params), { invalidateAll: true });
	}
</script>

<ActionPortal>
	<Action href={resolveRoute('/recipes/[recipeId]', $page.params)} title={$t('actions.back')}>
		<i class="icon-undo-2"></i>
	</Action>

	<Action on:click={deleteRecipe} title={$t('actions.recipe.delete')}>
		<i class="icon-trash"></i>
	</Action>

	<Action on:click={saveRecipe} title={$t('actions.save')}>
		<i class="icon-save"></i>
	</Action>
</ActionPortal>

<Editor bind:value={content} />

<BooleanModal bind:this={deleteModal} let:close>
	<p>{$t('recipe.modal.delete')}</p>

	<div class="flex justify-between mt-5">
		<button
			class="bg-gray-200 px-3 py-1 transition hover:(text-white bg-blue)"
			on:click={() => close(false)}
		>
			{$t('actions.no')}
		</button>
		<button
			class="bg-gray-200 px-3 py-1 transition hover:(text-white bg-red)"
			on:click={() => close(true)}
		>
			{$t('actions.yes')}
		</button>
	</div>
</BooleanModal>
