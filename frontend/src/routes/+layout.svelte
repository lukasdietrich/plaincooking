<script>
	import '@unocss/reset/tailwind-compat.css';

	import Header from './Header.svelte';
</script>

<div class="app">
	<Header />

	<main class="max-w-screen-lg mx-auto p-5">
		<slot />
	</main>
</div>
