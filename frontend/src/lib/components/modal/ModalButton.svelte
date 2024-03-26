<script lang="ts" generics="T">
	import { getClose } from './context';

	// eslint-disable-next-line no-undef
	let result: T;
	let cssClasses: string = '';

	export { cssClasses as class, result };

	// eslint-disable-next-line no-undef
	const close = getClose<T>();
</script>

<button
	on:click={() => close(result)}
	class="rounded-sm bg-gray-200 px-4 py-2 transition hover:(text-white bg-gray-500) {cssClasses}"
>
	<slot />
</button>
