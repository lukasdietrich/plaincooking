<script>
	import '@unocss/reset/tailwind.css';
	import 'lucide-static/font/lucide.css';

	import { page } from '$app/stores';

	import Header from './Header.svelte';

	$: title = $page.data['title'] ?? 'Plaincooking';
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

<div class="min-h-screen bg-white">
	<Header />

	<main class="max-w-screen-lg mx-auto p-5">
		<slot />
	</main>
</div>
