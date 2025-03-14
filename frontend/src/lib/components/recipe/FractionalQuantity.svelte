<script lang="ts">
	import type Fraction from 'fraction.js';

	export let quantity: Fraction;
	export let unit: string | undefined;

	$: wholeNumbers = quantity.n / quantity.d;
	$: numerator = quantity.n % quantity.d;
	$: denominator = quantity.d;
</script>

{#if wholeNumbers > 0}
	<span>
		{wholeNumbers}
	</span>
{/if}

{#if numerator !== 0n}
	<span class="relative text-sm">
		<span class="relative -top-1">{numerator}</span>
		<span class="-mx-0.75">/</span>
		<span class="relative top-0.5">{denominator}</span>
	</span>
{/if}

{#if unit}
	<span>
		{unit}
	</span>
{/if}
