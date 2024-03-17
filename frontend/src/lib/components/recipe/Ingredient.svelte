<script lang="ts">
	import type { Tokens } from '$lib/recipe';
	import FractionalQuantity from './FractionalQuantity.svelte';

	export let token: Tokens.Ingredient;

	const id = Math.random().toString(16);

	let done = false;
</script>

<div class="inline-block select-none bg-emerald-100 rounded-full my-1 px-1">
	<label for={id} class="font-semibold px-2 py-0.5 thickness-1" class:line-through={done}>
		{#if token.quantity}
			<FractionalQuantity quantity={token.quantity} unit={token.unit} />
		{/if}

		<span class:ml-2={token.quantity}>
			{token.ingredient}
		</span>
	</label>
	<input type="checkbox" class="mr-2 accent-emerald" {id} bind:checked={done} />
</div>
