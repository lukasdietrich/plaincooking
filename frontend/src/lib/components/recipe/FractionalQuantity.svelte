<script lang="ts">
	import type { Fraction } from 'mathjs';

	export let quantity: Fraction;
	export let unit: string | undefined;

	$: wholeNumbers = Math.floor(quantity.n / quantity.d);
	$: numerator = quantity.n % quantity.d;
	$: denominator = quantity.d;
</script>

{#if wholeNumbers > 0}
	<span>
		{wholeNumbers}
	</span>
{/if}

{#if numerator !== 0}
	<span class="text-sm">
		<sup>{numerator}</sup>/<sub>{denominator}</sub>
	</span>
{/if}

{#if unit}
	<span>
		{unit}
	</span>
{/if}
