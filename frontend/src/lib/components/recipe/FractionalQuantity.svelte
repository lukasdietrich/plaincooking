<script lang="ts">
	import type { Fraction } from 'mathjs';

	export let quantity: Fraction;
	export let unit: string | undefined;

	$: isFraction = quantity.d !== 1;
</script>

<span class:diagonal-fractions={isFraction}>
	{quantity.n}{#if isFraction}/{quantity.d}{/if}
</span>

{#if unit}
	<span>
		{unit}
	</span>
{/if}
