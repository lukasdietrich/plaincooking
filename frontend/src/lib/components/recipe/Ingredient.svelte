<script lang="ts">
	import type { Tokens } from '$lib/recipe';

	export let token: Tokens.Ingredient;

	const id = Math.random().toString(16);

	let done = false;

	$: isFraction = token.quantity?.d !== 1;
</script>

<div class="inline-block items-center select-none bg-emerald-100 rounded-full my-1 px-1">
	<label
		for={id}
		class="inline-flex font-semibold px-2 py-0.5 thickness-1"
		class:line-through={done}
	>
		{#if token.quantity}
			<span class:diagonal-fractions={isFraction}>
				{token.quantity.n}{#if isFraction}/{token.quantity.d}{/if}
			</span>

			{#if token.unit}
				<span>
					{token.unit}
				</span>
			{/if}
		{/if}

		<span class:ml-2={token.quantity}>
			{token.ingredient}
		</span>
	</label>
	<input type="checkbox" class="mr-2 accent-emerald" {id} bind:checked={done} />
</div>
